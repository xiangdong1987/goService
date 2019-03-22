package main

import (
	"context"
	"encoding/json"
	"github.com/go-kit/kit/endpoint"
	"net/http"
)

type matchRequest struct {
	S string `json:"s"`
}

type matchResponse struct {
	V     string `json:"text"`
	Level string `json:"level"`
	Err   string `json:"err,omitempty"` // errors don't JSON-marshal, so we use a string
}

type countRequest struct {
	S string `json:"s"`
}

type countResponse struct {
	V int `json:"v"`
}

func makeMatchEndpoint(svc FilterService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(matchRequest)
		level, v, err := svc.Match(req.S)
		if err != nil {
			return matchResponse{v, level, err.Error()}, nil
		}
		return matchResponse{v, level, ""}, nil
	}
}
func makeCountEndpoint(svc FilterService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(countRequest)
		v := svc.Count(req.S)
		return countResponse{v}, nil
	}
}

func decodeMatchRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request matchRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func decodeCountRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request countRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
