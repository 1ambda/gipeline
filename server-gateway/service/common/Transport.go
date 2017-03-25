package common

import (
	"encoding/json"
	"net/http"

	"golang.org/x/net/context"
)

type HasError interface {
	error() error
}

type ErrResponse struct {
	Error error `json:"error,omitempty"`
}

func (r ErrResponse) error() error {
	return r.Error
}

func NewErrResponse(err error) *ErrResponse {
	return &ErrResponse{Error: err}
}

func EncodeError(_ context.Context, err error, w http.ResponseWriter) {
	// TODO: switch errors

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}

func DecodeJsonEmptyRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req interface{}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}

	return req, nil
}

func DecodeEmptyRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req interface{}
	return req, nil
}

func EncodeResponse(ctx context.Context, w http.ResponseWriter, res interface{}) error {
	if e, hasError := res.(HasError); hasError && e.error() != nil {
		EncodeError(ctx, e.error(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(res)
}
