package transport

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	httpHandler "net/http"
	"strconv"

	"backend/internal/casestudy"

	model "backend/pkg/casestudy"

	"github.com/go-kit/kit/transport/http"
	"github.com/go-kit/log"
	"github.com/gorilla/mux"
)

func NewHTTPHandler(endpoints casestudy.Endpoints, logger log.Logger) httpHandler.Handler {
	r := mux.NewRouter()

	options := []http.ServerOption{
		http.ServerErrorLogger(logger),
		http.ServerErrorEncoder(encodeError),
	}

	r.Methods("POST").Path("/case").Handler(http.NewServer(
		endpoints.Create,
		decodeCreateRequest,
		encodeResponse,
	))

	r.Methods("GET").Path("/case-study/{id}").Handler(http.NewServer(
		endpoints.Get,
		decodeGetRequest,
		encodeResponse,
		options...,
	))

	r.Methods("GET").Path("/case-studies").Handler(http.NewServer(
		endpoints.GetAll,
		decodeGetAllRequest,
		encodeResponse,
		options...,
	))

	r.Methods("POST").Path("/upload").Handler(http.NewServer(
		endpoints.Upload,
		decodeUploadRequest,
		encodeResponse,
		options...,
	))

	return r
}

func decodeUploadRequest(_ context.Context, r *httpHandler.Request) (interface{}, error) {
	file, _, err := r.FormFile("file")
	if err != nil {
		return nil, err
	}
	fileName := r.FormValue("fileName")
	if fileName == "" {
		return nil, fmt.Errorf("fileName is required")
	}
	return struct {
		File     io.Reader
		FileName string
	}{
		File:     file,
		FileName: fileName,
	}, nil
}

func decodeCreateRequest(_ context.Context, r *httpHandler.Request) (interface{}, error) {
	var caseStudy model.CaseStudy
	if err := json.NewDecoder(r.Body).Decode(&caseStudy); err != nil {
		return nil, err
	}
	return &caseStudy, nil
}

func decodeGetRequest(_ context.Context, r *httpHandler.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		return nil, err
	}
	return casestudy.GetRequest{ID: uint(id)}, nil
}

func decodeGetAllRequest(_ context.Context, r *httpHandler.Request) (interface{}, error) {
	return nil, nil
}

func encodeResponse(ctx context.Context, w httpHandler.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func encodeError(_ context.Context, err error, w httpHandler.ResponseWriter) {
	w.WriteHeader(httpHandler.StatusInternalServerError)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}
