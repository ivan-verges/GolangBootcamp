package service

import (
	data "BairesDev/Level6/02_Final_Microservices/questions/data"
	model "BairesDev/Level6/02_Final_Microservices/questions/model"
	"context"
	"fmt"
)

type QuestionRESTService interface {
	GetAll(ctx context.Context) (*[]model.Question, error)
	GetById(ctx context.Context, qstId string) (*model.Question, error)
	Create(ctx context.Context, qst model.Question) (*model.Question, error)
	Update(ctx context.Context, qstId string, qst model.Question) (*model.Question, error)
	Delete(ctx context.Context, qstId string) (*model.Question, error)
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

func (s *service) GetAll(ctx context.Context) (*[]model.Question, error) {
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

func (s *service) GetById(ctx context.Context, qstId string) (*model.Question, error) {
	switch s.dbType {
	case "mysql":
		{
			mySqlDb := data.NewMySqlRepo(model.GetMySqlConnection())
			return mySqlDb.GetById(ctx, qstId)
		}
	case "redis":
		{
			redisDb := data.NewRedisRepo(model.GetRedisConnection())
			return redisDb.GetById(ctx, qstId)
		}
	}

	return nil, fmt.Errorf("wrong or missing db type")
}

func (s *service) Create(ctx context.Context, qst model.Question) (*model.Question, error) {
	switch s.dbType {
	case "mysql":
		{
			mySqlDb := data.NewMySqlRepo(model.GetMySqlConnection())
			return mySqlDb.Create(ctx, qst)
		}
	case "redis":
		{
			redisDb := data.NewRedisRepo(model.GetRedisConnection())
			return redisDb.Create(ctx, qst)
		}
	}

	return nil, fmt.Errorf("wrong or missing db type")
}

func (s *service) Update(ctx context.Context, qstId string, qst model.Question) (*model.Question, error) {
	switch s.dbType {
	case "mysql":
		{
			mySqlDb := data.NewMySqlRepo(model.GetMySqlConnection())
			return mySqlDb.Update(ctx, qstId, qst)
		}
	case "redis":
		{
			redisDb := data.NewRedisRepo(model.GetRedisConnection())
			return redisDb.Update(ctx, qstId, qst)
		}
	}

	return nil, fmt.Errorf("wrong or missing db type")
}

func (s *service) Delete(ctx context.Context, qstId string) (*model.Question, error) {
	switch s.dbType {
	case "mysql":
		{
			mySqlDb := data.NewMySqlRepo(model.GetMySqlConnection())
			return mySqlDb.Delete(ctx, qstId)
		}
	case "redis":
		{
			redisDb := data.NewRedisRepo(model.GetRedisConnection())
			return redisDb.Delete(ctx, qstId)
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
