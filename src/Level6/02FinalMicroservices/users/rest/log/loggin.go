package log

import (
	model "BairesDev/Level6/02_Final_Microservices/users/model"
	service "BairesDev/Level6/02_Final_Microservices/users/rest/service"
	"context"
	"time"

	"github.com/go-kit/log"
)

func NewLoggingMiddleware(logger log.Logger, next service.UserRESTService) logmw {
	return logmw{logger, next}
}

type logmw struct {
	logger log.Logger
	service.UserRESTService
}

func (mw logmw) GetAll(ctx context.Context) (users *[]model.User, err error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "GetAll",
			"input", "",
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	return mw.UserRESTService.GetAll(ctx)
}

func (mw logmw) GetById(ctx context.Context, usrId string) (user *model.User, err error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "GetById",
			"input", usrId,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	return mw.UserRESTService.GetById(ctx, usrId)
}

func (mw logmw) Create(ctx context.Context, usr model.User) (user *model.User, err error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "Create",
			"input", usr.UserName,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	return mw.UserRESTService.Create(ctx, usr)
}

func (mw logmw) Update(ctx context.Context, userId string, usr model.User) (user *model.User, err error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "Update",
			"input", userId,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	return mw.UserRESTService.Update(ctx, userId, usr)
}

func (mw logmw) Delete(ctx context.Context, userId string) (user *model.User, err error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "Delete",
			"input", userId,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	return mw.UserRESTService.Delete(ctx, userId)
}
