package grpc

import (
	data "BairesDev/Level6/02_Final_Microservices/answers/data"
	model "BairesDev/Level6/02_Final_Microservices/answers/grpc/model"
	rest "BairesDev/Level6/02_Final_Microservices/answers/model"
	"context"
	"fmt"
)

type AnswerGRPCService struct {
	dbType string
}

func NewGrpcAnswerService(dbType string) model.AnswersServer {
	return &AnswerGRPCService{dbType: dbType}
}

func (service *AnswerGRPCService) GetAnswer(ctx context.Context, input *model.Input) (*model.Answer, error) {
	switch service.dbType {
	case "mysql":
		{
			mySqlDb := data.NewMySqlRepo(rest.GetMySqlConnection())
			tmp, err := mySqlDb.GetById(ctx, input.AnswerId)
			if err != nil {
				return nil, err
			}

			result := model.Answer{ID: int64(tmp.ID), AnswerId: tmp.AnswerId, UserId: tmp.UserId, QuestionId: tmp.QuestionId, Answer: tmp.Answer}
			return &result, nil
		}
	case "redis":
		{
			redisDb := data.NewRedisRepo(rest.GetRedisConnection())
			tmp, err := redisDb.GetById(ctx, input.AnswerId)
			if err != nil {
				return nil, err
			}

			result := model.Answer{ID: int64(tmp.ID), AnswerId: tmp.AnswerId, UserId: tmp.UserId, QuestionId: tmp.QuestionId, Answer: tmp.Answer}
			return &result, nil
		}
	}

	return nil, fmt.Errorf("wrong or missing db type")
}

func (service *AnswerGRPCService) GetUserWithQuestionsAndAnswers(ctx context.Context, input *model.Input) (*model.UserQuestionsAnswers, error) {
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
