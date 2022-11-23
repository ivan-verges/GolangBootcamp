package log

import (
	model "BairesDev/Level6/02_Final_Microservices/questions/model"
	service "BairesDev/Level6/02_Final_Microservices/questions/rest/service"
	"context"
	"time"

	"github.com/go-kit/log"
)

func NewLoggingMiddleware(logger log.Logger, next service.QuestionRESTService) logmw {
	return logmw{logger, next}
}

type logmw struct {
	logger log.Logger
	service.QuestionRESTService
}

func (mw logmw) GetAll(ctx context.Context) (users *[]model.Question, err error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "GetAll",
			"input", "",
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	return mw.QuestionRESTService.GetAll(ctx)
}

func (mw logmw) GetById(ctx context.Context, qstId string) (user *model.Question, err error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "GetById",
			"input", qstId,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	return mw.QuestionRESTService.GetById(ctx, qstId)
}

func (mw logmw) Create(ctx context.Context, qst model.Question) (user *model.Question, err error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "Create",
			"input", qst.Question,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	return mw.QuestionRESTService.Create(ctx, qst)
}

func (mw logmw) Update(ctx context.Context, questionId string, qst model.Question) (user *model.Question, err error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "Update",
			"input", questionId,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	return mw.QuestionRESTService.Update(ctx, questionId, qst)
}

func (mw logmw) Delete(ctx context.Context, questionId string) (user *model.Question, err error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "Delete",
			"input", questionId,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	return mw.QuestionRESTService.Delete(ctx, questionId)
}
