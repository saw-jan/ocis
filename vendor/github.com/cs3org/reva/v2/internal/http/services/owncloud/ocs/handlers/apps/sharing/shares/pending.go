// Copyright 2018-2021 CERN
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// In applying this license, CERN does not waive the privileges and immunities
// granted to it by virtue of its status as an Intergovernmental Organization
// or submit itself to any jurisdiction.

package shares

import (
	"context"
	"fmt"
	"net/http"
	"path"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	gateway "github.com/cs3org/go-cs3apis/cs3/gateway/v1beta1"
	rpc "github.com/cs3org/go-cs3apis/cs3/rpc/v1beta1"
	collaboration "github.com/cs3org/go-cs3apis/cs3/sharing/collaboration/v1beta1"
	ocmv1beta1 "github.com/cs3org/go-cs3apis/cs3/sharing/ocm/v1beta1"
	provider "github.com/cs3org/go-cs3apis/cs3/storage/provider/v1beta1"
	"github.com/cs3org/reva/v2/internal/http/services/owncloud/ocs/response"
	"github.com/cs3org/reva/v2/pkg/appctx"
	"github.com/cs3org/reva/v2/pkg/conversions"
	"github.com/cs3org/reva/v2/pkg/utils"
	"github.com/go-chi/chi/v5"
	"github.com/pkg/errors"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
)

const (
	// shareidkey is the key user to obtain the id of the share to update. It is present in the request URL.
	shareidkey string = "shareid"
)

// AcceptReceivedShare handles Post Requests on /apps/files_sharing/api/v1/shares/{shareid}
func (h *Handler) AcceptReceivedShare(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	shareID := chi.URLParam(r, shareidkey)

	if h.isFederatedReceivedShare(r, shareID) {
		h.updateReceivedFederatedShare(w, r, shareID, false)
		return
	}

	client, err := h.getClient()
	if err != nil {
		response.WriteOCSError(w, r, response.MetaServerError.StatusCode, "error getting grpc gateway client", err)
		return
	}

	rs, ocsResponse := getReceivedShareFromID(ctx, client, shareID)
	if ocsResponse != nil {
		response.WriteOCSResponse(w, r, *ocsResponse, nil)
		return
	}

	sharedResource, ocsResponse := getSharedResource(ctx, client, rs.Share.Share.ResourceId)
	if ocsResponse != nil {
		response.WriteOCSResponse(w, r, *ocsResponse, nil)
		return
	}

	lrs, ocsResponse := getSharesList(ctx, client)
	if ocsResponse != nil {
		response.WriteOCSResponse(w, r, *ocsResponse, nil)
		return
	}

	// we need to sort the received shares by mount point in order to make things easier to evaluate.
	base := path.Base(sharedResource.GetInfo().GetPath())
	mount := base
	var mountedShares []*collaboration.ReceivedShare
	sharesToAccept := map[string]bool{shareID: true}
	for _, s := range lrs.Shares {
		if utils.ResourceIDEqual(s.Share.ResourceId, rs.Share.Share.GetResourceId()) {
			if s.State == collaboration.ShareState_SHARE_STATE_ACCEPTED {
				mount = s.MountPoint.Path
			} else {
				sharesToAccept[s.Share.Id.OpaqueId] = true
			}
		} else {
			if s.State == collaboration.ShareState_SHARE_STATE_ACCEPTED {
				s.Hidden = h.getReceivedShareHideFlagFromShareID(r.Context(), shareID)
				mountedShares = append(mountedShares, s)
			}
		}
	}

	compareMountPoint := func(i, j int) bool {
		return mountedShares[i].MountPoint.Path > mountedShares[j].MountPoint.Path
	}
	sort.Slice(mountedShares, compareMountPoint)

	// now we have a list of shares, we want to iterate over all of them and check for name collisions
	for i, ms := range mountedShares {
		if ms.MountPoint.Path == mount {
			// does the shared resource still exist?
			res, err := client.Stat(ctx, &provider.StatRequest{
				Ref: &provider.Reference{
					ResourceId: ms.Share.ResourceId,
				},
			})
			if err == nil && res.Status.Code == rpc.Code_CODE_OK {
				// The mount point really already exists, we need to insert a number into the filename
				ext := filepath.Ext(base)
				name := strings.TrimSuffix(base, ext)
				// be smart about .tar.(gz|bz) files
				if strings.HasSuffix(name, ".tar") {
					name = strings.TrimSuffix(name, ".tar")
					ext = ".tar" + ext
				}

				mount = fmt.Sprintf("%s (%s)%s", name, strconv.Itoa(i+1), ext)
			}
			// TODO we could delete shares here if the stat returns code NOT FOUND ... but listening for file deletes would be better
		}
	}
	// we need to add a path to the share
	receivedShare := &collaboration.ReceivedShare{
		Share: &collaboration.Share{
			Id: &collaboration.ShareId{OpaqueId: shareID},
		},
		State:  collaboration.ShareState_SHARE_STATE_ACCEPTED,
		Hidden: h.getReceivedShareHideFlagFromShareID(r.Context(), shareID),
		MountPoint: &provider.Reference{
			Path: mount,
		},
	}
	updateMask := &fieldmaskpb.FieldMask{Paths: []string{"state", "hidden", "mount_point"}}

	for id := range sharesToAccept {
		data := h.updateReceivedShare(w, r, receivedShare, updateMask)
		// only render the data for the changed share
		if id == shareID && data != nil {
			response.WriteOCSSuccess(w, r, []*conversions.ShareData{data})
		}
	}
}

// RejectReceivedShare handles DELETE Requests on /apps/files_sharing/api/v1/shares/{shareid}
func (h *Handler) RejectReceivedShare(w http.ResponseWriter, r *http.Request) {
	shareID := chi.URLParam(r, "shareid")

	if h.isFederatedReceivedShare(r, shareID) {
		h.updateReceivedFederatedShare(w, r, shareID, true)
		return
	}

	// we need to add a path to the share
	receivedShare := &collaboration.ReceivedShare{
		Share: &collaboration.Share{
			Id: &collaboration.ShareId{OpaqueId: shareID},
		},
		State:  collaboration.ShareState_SHARE_STATE_REJECTED,
		Hidden: h.getReceivedShareHideFlagFromShareID(r.Context(), shareID),
	}
	updateMask := &fieldmaskpb.FieldMask{Paths: []string{"state", "hidden"}}

	data := h.updateReceivedShare(w, r, receivedShare, updateMask)
	if data != nil {
		response.WriteOCSSuccess(w, r, []*conversions.ShareData{data})
	}
}

func (h *Handler) UpdateReceivedShare(w http.ResponseWriter, r *http.Request) {
	shareID := chi.URLParam(r, "shareid")
	hideFlag, _ := strconv.ParseBool(r.URL.Query().Get("hidden"))

	// unfortunately we need to get the share first to read the state
	client, err := h.getClient()
	if err != nil {
		response.WriteOCSError(w, r, response.MetaServerError.StatusCode, "error getting grpc gateway client", err)
		return
	}

	// we need to add a path to the share
	receivedShare := &collaboration.ReceivedShare{
		Share: &collaboration.Share{
			Id: &collaboration.ShareId{OpaqueId: shareID},
		},
		Hidden: hideFlag,
	}
	updateMask := &fieldmaskpb.FieldMask{Paths: []string{"state", "hidden"}}

	rs, _ := getReceivedShareFromID(r.Context(), client, shareID)
	if rs != nil && rs.Share != nil {
		receivedShare.State = rs.Share.State
	}

	data := h.updateReceivedShare(w, r, receivedShare, updateMask)
	if data != nil {
		response.WriteOCSSuccess(w, r, []*conversions.ShareData{data})
	}
	// TODO: do we need error handling here?
}

func (h *Handler) updateReceivedShare(w http.ResponseWriter, r *http.Request, receivedShare *collaboration.ReceivedShare, fieldMask *fieldmaskpb.FieldMask) *conversions.ShareData {
	ctx := r.Context()
	logger := appctx.GetLogger(ctx)

	updateShareRequest := &collaboration.UpdateReceivedShareRequest{
		Share:      receivedShare,
		UpdateMask: fieldMask,
	}

	client, err := h.getClient()
	if err != nil {
		response.WriteOCSError(w, r, response.MetaServerError.StatusCode, "error getting grpc gateway client", err)
		return nil
	}

	shareRes, err := client.UpdateReceivedShare(ctx, updateShareRequest)
	if err != nil {
		response.WriteOCSError(w, r, response.MetaServerError.StatusCode, "grpc update received share request failed", err)
		return nil
	}

	if shareRes.Status.Code != rpc.Code_CODE_OK {
		if shareRes.Status.Code == rpc.Code_CODE_NOT_FOUND {
			response.WriteOCSError(w, r, response.MetaNotFound.StatusCode, "not found", nil)
			return nil
		}
		response.WriteOCSError(w, r, response.MetaServerError.StatusCode, "grpc update received share request failed", errors.Errorf("code: %d, message: %s", shareRes.Status.Code, shareRes.Status.Message))
		return nil
	}

	rs := shareRes.GetShare()

	info, status, err := h.getResourceInfoByID(ctx, client, rs.Share.ResourceId)
	if err != nil || status.Code != rpc.Code_CODE_OK {
		h.logProblems(logger, status, err, "could not stat, skipping")
		response.WriteOCSError(w, r, response.MetaServerError.StatusCode, "grpc get resource info failed", errors.Errorf("code: %d, message: %s", status.Code, status.Message))
		return nil
	}

	data, err := conversions.CS3Share2ShareData(r.Context(), rs.Share)
	if err != nil {
		logger.Debug().Interface("share", rs.Share).Interface("shareData", data).Err(err).Msg("could not CS3Share2ShareData, skipping")
	}

	data.State = mapState(rs.GetState())
	data.Hidden = rs.GetHidden()

	h.addFileInfo(ctx, data, info)
	h.mapUserIds(r.Context(), client, data)

	if data.State == ocsStateAccepted {
		// Needed because received shares can be jailed in a folder in the users home
		data.Path = path.Join(h.sharePrefix, path.Base(info.Path))
	}

	return data
}

func (h *Handler) updateReceivedFederatedShare(w http.ResponseWriter, r *http.Request, shareID string, rejectShare bool) {
	ctx := r.Context()

	client, err := h.getClient()
	if err != nil {
		response.WriteOCSError(w, r, response.MetaServerError.StatusCode, "error getting grpc gateway client", err)
		return
	}

	share, err := client.GetReceivedOCMShare(ctx, &ocmv1beta1.GetReceivedOCMShareRequest{
		Ref: &ocmv1beta1.ShareReference{
			Spec: &ocmv1beta1.ShareReference_Id{
				Id: &ocmv1beta1.ShareId{
					OpaqueId: shareID,
				},
			},
		},
	})
	if err != nil {
		response.WriteOCSError(w, r, response.MetaServerError.StatusCode, "grpc update received share request failed", err)
		return
	}
	if share.Status.Code != rpc.Code_CODE_OK {
		if share.Status.Code == rpc.Code_CODE_NOT_FOUND {
			response.WriteOCSError(w, r, response.MetaNotFound.StatusCode, "not found", nil)
			return
		}
		response.WriteOCSError(w, r, response.MetaServerError.StatusCode, "grpc update received share request failed", errors.Errorf("code: %d, message: %s", share.Status.Code, share.Status.Message))
		return
	}

	req := &ocmv1beta1.UpdateReceivedOCMShareRequest{
		Share: &ocmv1beta1.ReceivedShare{
			Id: &ocmv1beta1.ShareId{
				OpaqueId: shareID,
			},
		},
		UpdateMask: &fieldmaskpb.FieldMask{Paths: []string{"state"}},
	}
	if rejectShare {
		req.Share.State = ocmv1beta1.ShareState_SHARE_STATE_REJECTED
	} else {
		req.Share.State = ocmv1beta1.ShareState_SHARE_STATE_ACCEPTED
	}

	updateRes, err := client.UpdateReceivedOCMShare(ctx, req)
	if err != nil {
		response.WriteOCSError(w, r, response.MetaServerError.StatusCode, "grpc update received share request failed", err)
		return
	}

	if updateRes.Status.Code != rpc.Code_CODE_OK {
		if updateRes.Status.Code == rpc.Code_CODE_NOT_FOUND {
			response.WriteOCSError(w, r, response.MetaNotFound.StatusCode, "not found", nil)
			return
		}
		response.WriteOCSError(w, r, response.MetaServerError.StatusCode, "grpc update received share request failed", errors.Errorf("code: %d, message: %s", updateRes.Status.Code, updateRes.Status.Message))
		return
	}

	data, err := conversions.ReceivedOCMShare2ShareData(share.Share, h.ocmLocalMount(share.Share))
	if err != nil {
		response.WriteOCSError(w, r, response.MetaServerError.StatusCode, "grpc update received share request failed", err)
		return
	}
	h.mapUserIdsReceivedFederatedShare(ctx, client, data)
	data.State = mapOCMState(req.Share.State)
	response.WriteOCSSuccess(w, r, []*conversions.ShareData{data})
}

// getReceivedShareHideFlagFromShareId returns the hide flag of a received share based on its ID.
func (h *Handler) getReceivedShareHideFlagFromShareID(ctx context.Context, shareID string) bool {
	client, err := h.getClient()
	if err != nil {
		return false
	}
	rs, _ := getReceivedShareFromID(ctx, client, shareID)
	if rs != nil {
		return rs.GetShare().GetHidden()
	}
	return false
}

// getReceivedShareFromID uses a client to the gateway to fetch a share based on its ID.
func getReceivedShareFromID(ctx context.Context, client gateway.GatewayAPIClient, shareID string) (*collaboration.GetReceivedShareResponse, *response.Response) {
	s, err := client.GetReceivedShare(ctx, &collaboration.GetReceivedShareRequest{
		Ref: &collaboration.ShareReference{
			Spec: &collaboration.ShareReference_Id{
				Id: &collaboration.ShareId{
					OpaqueId: shareID,
				}},
		},
	})

	if err != nil {
		e := errors.Wrap(err, fmt.Sprintf("could not get share with ID: `%s`", shareID))
		return nil, arbitraryOcsResponse(response.MetaServerError.StatusCode, e.Error())
	}

	if s.Status.Code != rpc.Code_CODE_OK {
		if s.Status.Code == rpc.Code_CODE_NOT_FOUND {
			e := fmt.Errorf("share not found")
			return nil, arbitraryOcsResponse(response.MetaNotFound.StatusCode, e.Error())
		}

		e := fmt.Errorf("invalid share: %s", s.GetStatus().GetMessage())
		return nil, arbitraryOcsResponse(response.MetaBadRequest.StatusCode, e.Error())
	}

	return s, nil
}

// getSharedResource attempts to get a shared resource from the storage from the resource reference.
func getSharedResource(ctx context.Context, client gateway.GatewayAPIClient, resID *provider.ResourceId) (*provider.StatResponse, *response.Response) {
	res, err := client.Stat(ctx, &provider.StatRequest{
		Ref: &provider.Reference{
			ResourceId: resID,
		},
	})
	if err != nil {
		e := fmt.Errorf("could not get reference")
		return nil, arbitraryOcsResponse(response.MetaServerError.StatusCode, e.Error())
	}

	if res.Status.Code != rpc.Code_CODE_OK {
		if res.Status.Code == rpc.Code_CODE_NOT_FOUND {
			e := fmt.Errorf("not found")
			return nil, arbitraryOcsResponse(response.MetaNotFound.StatusCode, e.Error())
		}
		e := fmt.Errorf(res.GetStatus().GetMessage())
		return nil, arbitraryOcsResponse(response.MetaServerError.StatusCode, e.Error())
	}

	return res, nil
}

// getSharedResource gets the list of all shares for the current user.
func getSharesList(ctx context.Context, client gateway.GatewayAPIClient) (*collaboration.ListReceivedSharesResponse, *response.Response) {
	shares, err := client.ListReceivedShares(ctx, &collaboration.ListReceivedSharesRequest{})
	if err != nil {
		e := errors.Wrap(err, "error getting shares list")
		return nil, arbitraryOcsResponse(response.MetaNotFound.StatusCode, e.Error())
	}

	if shares.Status.Code != rpc.Code_CODE_OK {
		if shares.Status.Code == rpc.Code_CODE_NOT_FOUND {
			e := fmt.Errorf("not found")
			return nil, arbitraryOcsResponse(response.MetaNotFound.StatusCode, e.Error())
		}
		e := fmt.Errorf(shares.GetStatus().GetMessage())
		return nil, arbitraryOcsResponse(response.MetaServerError.StatusCode, e.Error())
	}
	return shares, nil
}

// arbitraryOcsResponse abstracts the boilerplate that is creating a response.Response struct.
func arbitraryOcsResponse(statusCode int, message string) *response.Response {
	r := response.Response{
		OCS: &response.Payload{
			XMLName: struct{}{},
			Meta:    response.Meta{},
			Data:    nil,
		},
	}

	r.OCS.Meta.StatusCode = statusCode
	r.OCS.Meta.Message = message
	return &r
}
