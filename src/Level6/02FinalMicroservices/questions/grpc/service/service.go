package grpc

import (
	data "BairesDev/Level6/02_Final_Microservices/questions/data"
	model "BairesDev/Level6/02_Final_Microservices/questions/grpc/model"
	rest "BairesDev/Level6/02_Final_Microservices/questions/model"
	"context"
	"fmt"
)

type QuestionGRPCService struct {
	dbType string
}

func NewGrpcUserService(dbType string) model.QuestionsServer {
	return &QuestionGRPCService{dbType: dbType}
}

func (service *QuestionGRPCService) GetQuestion(ctx context.Context, input *model.Input) (*model.Question, error) {
	switch service.dbType {
	case "mysql":
		{
			mySqlDb := data.NewMySqlRepo(rest.GetMySqlConnection())
			tmp, err := mySqlDb.GetById(ctx, input.QuestionId)
			if err != nil {
				return nil, err
			}

			result := model.Question{ID: int64(tmp.ID), QuestionId: tmp.QuestionId, UserId: tmp.UserId, Question: tmp.Question}
			return &result, nil
		}
	case "redis":
		{
			redisDb := data.NewRedisRepo(rest.GetRedisConnection())
			tmp, err := redisDb.GetById(ctx, input.QuestionId)
			if err != nil {
				return nil, err
			}

			result := model.Question{ID: int64(tmp.ID), QuestionId: tmp.QuestionId, UserId: tmp.UserId, Question: tmp.Question}
			return &result, nil
		}
	}

	return nil, fmt.Errorf("wrong or missing db type")
}

func (service *QuestionGRPCService) GetUserWithQuestionsAndAnswers(ctx context.Context, input *model.Input) (*model.UserQuestionsAnswers, error) {
	switch service.dbType {
	case "mysql":
		{
			mySqlDb := data.NewMySqlRepo(rest.GetMySqlConnection())
			tmp := mySqlDb.GetUserWithQuestionsAndAnswers(ctx, input.UserId)
			if tmp == nil {
				return nil, fmt.Errorf("Error Getting User with Questions and Answers")
			}

			result := model.UserQuestionsAnswers{UserId: tmp.UserId, UserName: tmp.UserName}
			for _, qa := range tmp.QuestionsWithAnswers {
				result.QuestionsWithAnswers = append(result.QuestionsWithAnswers, &model.QuestionAnswer{QuestionId: qa.QuestionId, UserId: qa.UserId, Question: qa.Question, Answer: &model.Answer{ID: int64(qa.Answer.ID), AnswerId: qa.Answer.AnswerId, UserId: qa.Answer.UserId, QuestionId: qa.Answer.QuestionId, Answer: qa.Answer.Answer}})
			}

			return &result, nil
		}
	case "redis":
		{
			redisDb := data.NewRedisRepo(rest.GetRedisConnection())
			tmp := redisDb.GetUserWithQuestionsAndAnswers(ctx, input.UserId)
			if tmp == nil {
				return nil, fmt.Errorf("Error Getting User with Questions and Answers")
			}

			result := model.UserQuestionsAnswers{UserId: tmp.UserId, UserName: tmp.UserName}
			for _, qa := range tmp.QuestionsWithAnswers {
				result.QuestionsWithAnswers = append(result.QuestionsWithAnswers, &model.QuestionAnswer{QuestionId: qa.QuestionId, UserId: qa.UserId, Question: qa.Question, Answer: &model.Answer{ID: int64(qa.Answer.ID), AnswerId: qa.Answer.AnswerId, UserId: qa.Answer.UserId, QuestionId: qa.Answer.QuestionId, Answer: qa.Answer.Answer}})
			}

			return &result, nil
		}
	}

	return nil, fmt.Errorf("wrong or missing db type")
}
