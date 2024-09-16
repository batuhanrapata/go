package casestudy

import (
	firebase "backend/internal/firebase"
	"backend/pkg/casestudy"
	"context"
	"fmt"
	"io"

	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	Create endpoint.Endpoint
	Get    endpoint.Endpoint
	GetAll endpoint.Endpoint
	Upload endpoint.Endpoint
}

func MakeEndpoints(s Service) Endpoints {
	return Endpoints{
		Create: makeCreateEndpoint(s),
		Get:    makeGetEndpoint(s),
		GetAll: makeGetAllEndpoint(s),
		Upload: MakeUploadEndpoint(s),
	}
}

func makeCreateEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*casestudy.CaseStudy)
		err := s.Create(ctx, req)
		return nil, err
	}
}

func makeGetEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetRequest)
		cs, err := s.Get(ctx, req.ID)
		return cs, err
	}
}

func makeGetAllEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		caseStudies, err := s.GetAll(ctx)
		return caseStudies, err
	}
}

func MakeUploadEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(struct {
			File     io.Reader
			FileName string
		})
		if !ok {
			return nil, fmt.Errorf("invalid request")
		}

		url, err := firebase.UploadImage(req.File, req.FileName)
		if err != nil {
			return nil, err
		}

		return struct {
			URL string `json:"url"`
		}{URL: url}, nil
	}
}

type GetRequest struct {
	ID uint
}
