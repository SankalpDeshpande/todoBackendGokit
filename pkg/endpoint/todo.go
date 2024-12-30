package endpoint

import (
	"context"
	"gokit-crud-app/pkg/service"

	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
    Create endpoint.Endpoint
    Get    endpoint.Endpoint
    Update endpoint.Endpoint
    Delete endpoint.Endpoint
	GetAll endpoint.Endpoint
}

func MakeEndpoints(s service.TodoService) Endpoints {
    return Endpoints{
        Create: makeCreateEndpoint(s),
        Get:    makeGetEndpoint(s),
        Update: makeUpdateEndpoint(s),
        Delete: makeDeleteEndpoint(s),
		GetAll: makeGetAllEndpoint(s),
    }
}

func makeCreateEndpoint(s service.TodoService) endpoint.Endpoint {
    return func(ctx context.Context, request interface{}) (interface{}, error) {
        req := request.(service.Todo)
        id, err := s.Create(ctx, req)
        return map[string]interface{}{"id": id}, err
    }
}

func makeGetEndpoint(s service.TodoService) endpoint.Endpoint {
    return func(ctx context.Context, request interface{}) (interface{}, error) {
        id := request.(string)
        todo, err := s.Get(ctx, id)
        return todo, err
    }
}

func makeGetAllEndpoint(s service.TodoService) endpoint.Endpoint {
    return func(ctx context.Context, request interface{}) (interface{}, error) {
        todos, err := s.GetAll(ctx)
        return todos, err
    }
}

func makeUpdateEndpoint(s service.TodoService) endpoint.Endpoint {
    return func(ctx context.Context, request interface{}) (interface{}, error) {
        req := request.(service.Todo)
        err := s.Update(ctx, req)
        return nil, err
    }
}

func makeDeleteEndpoint(s service.TodoService) endpoint.Endpoint {
    return func(ctx context.Context, request interface{}) (interface{}, error) {
        id := request.(string)
        err := s.Delete(ctx, id)
        return nil, err
    }
}
