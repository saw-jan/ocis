package service

import (
	"context"
//	"fmt"
//	"io/ioutil"
	"os"
	"path/filepath"

//	merrors "github.com/asim/go-micro/v3/errors"
	"github.com/blevesearch/bleve/v2"
	"github.com/blevesearch/bleve/v2/analysis/analyzer/keyword"
	"github.com/owncloud/ocis/ocis-pkg/log"
	"github.com/owncloud/ocis/search/pkg/config"
	"github.com/owncloud/ocis/search/pkg/proto/v0"
//	"google.golang.org/protobuf/encoding/protojson"
)

type BleveDocument struct {
	Metadata map[string]*proto.Term `json:"metadata"`
}

// New returns a new instance of Service
func New(opts ...Option) (s *Service, err error) {
	options := newOptions(opts...)
	logger := options.Logger
	cfg := options.Config

	indexMapping := bleve.NewIndexMapping()
	// keep all symbols in terms to allow exact matching, eg. emails
	indexMapping.DefaultAnalyzer = keyword.Name

	s = &Service{
		id:     cfg.Service.Namespace + "." + cfg.Service.Name,
		log:    logger,
		Config: cfg,
	}

	indexDir := filepath.Join(cfg.Datapath, "index.bleve")
	// for now recreate index on every start
	if err = os.RemoveAll(indexDir); err != nil {
		return nil, err
	}
	if s.index, err = bleve.New(indexDir, indexMapping); err != nil {
		return
	}
	// if err = s.indexRecords(recordsDir); err != nil {
	//	return nil, err
	// }
	return
}

// Service implements the AccountsServiceHandler interface
type Service struct {
        id     string
        log    log.Logger
        Config *config.Config
        index  bleve.Index
}

func (s *Service) Search(c context.Context, rreq *proto.SearchRequest, rres *proto.SearchResponse) error {
	return nil
}

func (s *Service) Index(c context.Context, rreq *proto.IndexRequest, rres *proto.IndexResponse) error {
	return nil
}

