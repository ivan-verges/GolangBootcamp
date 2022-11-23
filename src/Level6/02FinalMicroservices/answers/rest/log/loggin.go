package log

import (
	model "BairesDev/Level6/02_Final_Microservices/answers/model"
	service "BairesDev/Level6/02_Final_Microservices/answers/rest/service"
	"context"
	"time"

	"github.com/go-kit/log"
)

func NewLoggingMiddleware(logger log.Logger, next service.AnswerRESTService) logmw {
	return logmw{logger, next}
}

type logmw struct {
	logger log.Logger
	service.AnswerRESTService
}

func (mw logmw) GetAll(ctx context.Context) (answers *[]model.Answer, err error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "GetAll",
			"input", "",
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	return mw.AnswerRESTService.GetAll(ctx)
}

func (mw logmw) GetById(ctx context.Context, usrId string) (answer *model.Answer, err error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "GetById",
			"input", usrId,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	return mw.AnswerRESTService.GetById(ctx, usrId)
}

func (mw logmw) Create(ctx context.Context, qst model.Answer) (answer *model.Answer, err error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "Create",
			"input", qst.Answer,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	return mw.AnswerRESTService.Create(ctx, qst)
}

func (mw logmw) Update(ctx context.Context, answerId string, qst model.Answer) (answer *model.Answer, err error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "Update",
			"input", answerId,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	return mw.AnswerRESTService.Update(ctx, answerId, qst)
}

func (mw logmw) Delete(ctx context.Context, answerId string) (answer *model.Answer, err error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "Delete",
			"input", answerId,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	return mw.AnswerRESTService.Delete(ctx, answerId)
}
