package http

import (
	"context"

	model "BairesDev/Level6/02_Final_Microservices/answers/model"
	service "BairesDev/Level6/02_Final_Microservices/answers/rest/service"

	"github.com/go-kit/kit/endpoint"
)

func MakeGetAll(svc service.AnswerRESTService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		return svc.GetAll(ctx)
	}
}

func MakeGetById(svc service.AnswerRESTService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getByIdRequest)
		return svc.GetById(ctx, req.AnswerID)
	}
}

func MakeCreate(svc service.AnswerRESTService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(createRequest)
		return svc.Create(ctx, model.Answer{UserId: req.UserID, QuestionId: req.QuestionID, Answer: req.Answer})
	}
}

func MakeUpdate(svc service.AnswerRESTService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(updateRequest)
		return svc.Update(ctx, req.AnswerID, model.Answer{AnswerId: req.AnswerID, QuestionId: req.QuestionID, UserId: req.UserID, Answer: req.Answer})
	}
}

func MakeDelete(svc service.AnswerRESTService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(deleteRequest)
		return svc.Delete(ctx, req.AnswerID)
	}
}
