/*
Libre Graph API

Libre Graph is a free API for cloud collaboration inspired by the MS Graph API.

API version: v1.0.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package libregraph

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// DrivesRootApiService DrivesRootApi service
type DrivesRootApiService service

type ApiGetRootRequest struct {
	ctx        context.Context
	ApiService *DrivesRootApiService
	driveId    string
}

func (r ApiGetRootRequest) Execute() (*DriveItem, *http.Response, error) {
	return r.ApiService.GetRootExecute(r)
}

/*
GetRoot Get root from arbitrary space

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param driveId key: id of drive
 @return ApiGetRootRequest
*/
func (a *DrivesRootApiService) GetRoot(ctx context.Context, driveId string) ApiGetRootRequest {
	return ApiGetRootRequest{
		ApiService: a,
		ctx:        ctx,
		driveId:    driveId,
	}
}

// Execute executes the request
//  @return DriveItem
func (a *DrivesRootApiService) GetRootExecute(r ApiGetRootRequest) (*DriveItem, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *DriveItem
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DrivesRootApiService.GetRoot")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1.0/drives/{drive-id}/root"
	localVarPath = strings.Replace(localVarPath, "{"+"drive-id"+"}", url.PathEscape(parameterValueToString(r.driveId, "driveId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v OdataError
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
