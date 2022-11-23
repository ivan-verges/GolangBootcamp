package http

import (
	"context"

	model "BairesDev/Level6/02_Final_Microservices/questions/model"
	service "BairesDev/Level6/02_Final_Microservices/questions/rest/service"

	"github.com/go-kit/kit/endpoint"
)

func MakeGetAll(svc service.QuestionRESTService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		return svc.GetAll(ctx)
	}
}

func MakeGetById(svc service.QuestionRESTService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getByIdRequest)
		return svc.GetById(ctx, req.QuestionID)
	}
}

func MakeCreate(svc service.QuestionRESTService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(createRequest)
		return svc.Create(ctx, model.Question{UserId: req.UserID, Question: req.Question})
	}
}

func MakeUpdate(svc service.QuestionRESTService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(updateRequest)
		return svc.Update(ctx, req.QuestionID, model.Question{QuestionId: req.QuestionID, UserId: req.UserID, Question: req.Question})
	}
}

func MakeDelete(svc service.QuestionRESTService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(deleteRequest)
		return svc.Delete(ctx, req.QuestionID)
	}
}
