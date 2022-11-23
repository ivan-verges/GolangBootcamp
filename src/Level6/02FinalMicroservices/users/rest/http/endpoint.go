package http

import (
	"context"

	model "BairesDev/Level6/02_Final_Microservices/users/model"
	service "BairesDev/Level6/02_Final_Microservices/users/rest/service"

	"github.com/go-kit/kit/endpoint"
)

func MakeGetAll(svc service.UserRESTService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		return svc.GetAll(ctx)
	}
}

func MakeGetById(svc service.UserRESTService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getByIdRequest)
		return svc.GetById(ctx, req.UserID)
	}
}

func MakeCreate(svc service.UserRESTService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(createRequest)
		return svc.Create(ctx, model.User{UserName: req.UserName})
	}
}

func MakeUpdate(svc service.UserRESTService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(updateRequest)
		return svc.Update(ctx, req.UserID, model.User{UserId: req.UserID, UserName: req.UserName})
	}
}

func MakeDelete(svc service.UserRESTService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(deleteRequest)
		return svc.Delete(ctx, req.UserID)
	}
}
