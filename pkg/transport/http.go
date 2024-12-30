package transport

import (
	"context"
	"encoding/json"
	"errors"
	"gokit-crud-app/pkg/endpoint"
	"gokit-crud-app/pkg/service"
	"net/http"
	"strings"

	httptransport "github.com/go-kit/kit/transport/http"
)

func NewHTTPHandler(endpoints endpoint.Endpoints) http.Handler {
    mux := http.NewServeMux()
	
    mux.Handle("/tasks/create",    httptransport.NewServer(
        endpoints.Create,
        decodeCreateRequest,
        encodeResponse,
    ))
    mux.Handle("/tasks/get/", httptransport.NewServer(
        endpoints.Get,
        decodeGetRequest,
        encodeResponse,
    ))
    mux.Handle("/tasks/update", httptransport.NewServer(
        endpoints.Update,
        decodeUpdateRequest,
        encodeResponse,
    ))
    mux.Handle("/tasks/delete/", httptransport.NewServer(
        endpoints.Delete,
        decodeDeleteRequest,
        encodeResponse,
    ))
	mux.Handle("/tasks/getall", httptransport.NewServer(
        endpoints.GetAll,
        decodeGetAllRequest,
        encodeResponse,
    ))

    return mux
}

func decodeCreateRequest(_ context.Context, r *http.Request) (interface{}, error) {
    var req service.Todo
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        return nil, err
    }
    return req, nil
}

func decodeGetRequest(_ context.Context, r *http.Request) (interface{}, error) {
    // Extract the id from the URL path, e.g., /get/:id
    id := strings.TrimPrefix(r.URL.Path, "/tasks/get/")
    if id == "" {
        return nil, errors.New("id is required")
    }
    return id, nil
}

func decodeGetAllRequest(_ context.Context, r *http.Request) (interface{}, error) {
    return nil, nil
}


func decodeUpdateRequest(_ context.Context, r *http.Request) (interface{}, error) {
    var req service.Todo
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        return nil, err
    }
    return req, nil
}

func decodeDeleteRequest(_ context.Context, r *http.Request) (interface{}, error) {
    // Extract the id from the URL path, e.g., /get/:id
    id := strings.TrimPrefix(r.URL.Path, "/tasks/delete/")
    if id == "" {
        return nil, errors.New("id is required")
    }
    return id, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
    return json.NewEncoder(w).Encode(response)
}
