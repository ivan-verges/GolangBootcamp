package service

import (
	data "BairesDev/Level6/02_Final_Microservices/answers/data"
	model "BairesDev/Level6/02_Final_Microservices/answers/model"
	"context"
	"fmt"
)

type AnswerRESTService interface {
	GetAll(ctx context.Context) (*[]model.Answer, error)
	GetById(ctx context.Context, ansId string) (*model.Answer, error)
	Create(ctx context.Context, ans model.Answer) (*model.Answer, error)
	Update(ctx context.Context, ansId string, ans model.Answer) (*model.Answer, error)
	Delete(ctx context.Context, ansId string) (*model.Answer, error)
	Migrate(ctx context.Context) error
	GetAnswerByQuestionId(ctx context.Context, questionId string) *model.Answer
	GetQuestionsAnswersByUserId(ctx context.Context, userId string) *[]model.QuestionAnswer
	GetUserWithQuestionsAndAnswers(ctx context.Context, userId string) *model.UserQuestionsAnswers
}

type service struct {
	dbType string
}

func NewService(dbType string) *service {
	return &service{dbType: dbType}
}

func (s *service) GetAll(ctx context.Context) (*[]model.Answer, error) {
	switch s.dbType {
	case "mysql":
		{
			mySqlDb := data.NewMySqlRepo(model.GetMySqlConnection())
			return mySqlDb.GetAll(ctx)
		}
	case "redis":
		{
			redisDb := data.NewRedisRepo(model.GetRedisConnection())
			return redisDb.GetAll(ctx)
		}
	}

	return nil, fmt.Errorf("wrong or missing db type")
}

func (s *service) GetById(ctx context.Context, ansId string) (*model.Answer, error) {
	switch s.dbType {
	case "mysql":
		{
			mySqlDb := data.NewMySqlRepo(model.GetMySqlConnection())
			return mySqlDb.GetById(ctx, ansId)
		}
	case "redis":
		{
			redisDb := data.NewRedisRepo(model.GetRedisConnection())
			return redisDb.GetById(ctx, ansId)
		}
	}

	return nil, fmt.Errorf("wrong or missing db type")
}

func (s *service) Create(ctx context.Context, ans model.Answer) (*model.Answer, error) {
	switch s.dbType {
	case "mysql":
		{
			mySqlDb := data.NewMySqlRepo(model.GetMySqlConnection())
			return mySqlDb.Create(ctx, ans)
		}
	case "redis":
		{
			redisDb := data.NewRedisRepo(model.GetRedisConnection())
			return redisDb.Create(ctx, ans)
		}
	}

	return nil, fmt.Errorf("wrong or missing db type")
}

func (s *service) Update(ctx context.Context, ansId string, ans model.Answer) (*model.Answer, error) {
	switch s.dbType {
	case "mysql":
		{
			mySqlDb := data.NewMySqlRepo(model.GetMySqlConnection())
			return mySqlDb.Update(ctx, ansId, ans)
		}
	case "redis":
		{
			redisDb := data.NewRedisRepo(model.GetRedisConnection())
			return redisDb.Update(ctx, ansId, ans)
		}
	}

	return nil, fmt.Errorf("wrong or missing db type")
}

func (s *service) Delete(ctx context.Context, ansId string) (*model.Answer, error) {
	switch s.dbType {
	case "mysql":
		{
			mySqlDb := data.NewMySqlRepo(model.GetMySqlConnection())
			return mySqlDb.Delete(ctx, ansId)
		}
	case "redis":
		{
			redisDb := data.NewRedisRepo(model.GetRedisConnection())
			return redisDb.Delete(ctx, ansId)
		}
	}

	return nil, fmt.Errorf("wrong or missing db type")
}

func (s *service) Migrate(ctx context.Context) error {
	switch s.dbType {
	case "mysql":
		{
			mySqlDb := data.NewMySqlRepo(model.GetMySqlConnection())
			return mySqlDb.Migrate(ctx)
		}
	case "redis":
		{
			redisDb := data.NewRedisRepo(model.GetRedisConnection())
			return redisDb.Migrate(ctx)
		}
	}

	return fmt.Errorf("wrong or missing db type")
}

func (s *service) GetAnswerByQuestionId(ctx context.Context, questionId string) *model.Answer {
	switch s.dbType {
	case "mysql":
		{
			mySqlDb := data.NewMySqlRepo(model.GetMySqlConnection())
			return mySqlDb.GetAnswerByQuestionId(ctx, questionId)
		}
	case "redis":
		{
			redisDb := data.NewRedisRepo(model.GetRedisConnection())
			return redisDb.GetAnswerByQuestionId(ctx, questionId)
		}
	}

	return nil
}

func (s *service) GetQuestionsAnswersByUserId(ctx context.Context, userId string) *[]model.QuestionAnswer {
	switch s.dbType {
	case "mysql":
		{
			mySqlDb := data.NewMySqlRepo(model.GetMySqlConnection())
			return mySqlDb.GetQuestionsAnswersByUserId(ctx, userId)
		}
	case "redis":
		{
			redisDb := data.NewRedisRepo(model.GetRedisConnection())
			return redisDb.GetQuestionsAnswersByUserId(ctx, userId)
		}
	}

	return nil
}

func (s *service) GetUserWithQuestionsAndAnswers(ctx context.Context, userId string) *model.UserQuestionsAnswers {
	switch s.dbType {
	case "mysql":
		{
			mySqlDb := data.NewMySqlRepo(model.GetMySqlConnection())
			return mySqlDb.GetUserWithQuestionsAndAnswers(ctx, userId)
		}
	case "redis":
		{
			redisDb := data.NewRedisRepo(model.GetRedisConnection())
			return redisDb.GetUserWithQuestionsAndAnswers(ctx, userId)
		}
	}

	return nil
}
