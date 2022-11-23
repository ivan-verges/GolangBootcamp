package data

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	model "BairesDev/Level6/02_Final_Microservices/questions/model"

	"github.com/go-redis/redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type redisRepo struct {
	DB *redis.Client
}

func NewRedisRepo(connection model.RedisConnectionString) *redisRepo {
	return &redisRepo{
		DB: redis.NewClient(&redis.Options{Addr: connection.GetAddress(), Password: connection.Password, DB: connection.Instance}),
	}
}

type mySqlRepo struct {
	DB *gorm.DB
}

func NewMySqlRepo(connection model.SqlConnectionString) *mySqlRepo {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       connection.GetShortConnectionString(),
		DefaultStringSize:         256,
		DisableDatetimePrecision:  true,
		DontSupportRenameIndex:    true,
		DontSupportRenameColumn:   true,
		SkipInitializeWithVersion: false,
	}), &gorm.Config{})
	if err != nil {
		fmt.Println(fmt.Sprintf("Error Connecting to:%s", connection.GetShortConnectionString()))
		panic(err)
	}

	return &mySqlRepo{DB: db}
}

func (repo *redisRepo) GetAll(ctx context.Context) (*[]model.Question, error) {
	result := []model.Question{}
	for _, key := range repo.DB.Keys(ctx, "QST*").Val() {
		tmp := model.Question{}

		data, err := repo.DB.Get(ctx, key).Result()
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal([]byte(data), &tmp)
		if err != nil {
			return nil, err
		} else {
			tmp.QuestionId = strings.Replace(tmp.QuestionId, "QST", "", -1)
			result = append(result, tmp)
		}
	}

	return &result, nil
}

func (repo *mySqlRepo) GetAll(ctx context.Context) (*[]model.Question, error) {
	result := []model.Question{}
	repo.DB.Find(&result)

	for index := range result {
		result[index].QuestionId = strconv.FormatUint(uint64(result[index].ID), 10)
	}

	return &result, nil
}

func (repo *redisRepo) GetById(ctx context.Context, qstId string) (*model.Question, error) {
	if !strings.Contains(qstId, "QST") {
		qstId = "QST" + qstId
	}

	data, err := repo.DB.Get(ctx, qstId).Result()
	if err != nil {
		return nil, err
	}

	result := model.Question{}
	err = json.Unmarshal([]byte(data), &result)
	if err != nil {
		return nil, err
	}

	result.QuestionId = strings.Replace(result.QuestionId, "QST", "", -1)

	return &result, nil
}

func (repo *mySqlRepo) GetById(ctx context.Context, qstId string) (*model.Question, error) {
	result := model.Question{}
	id, err := strconv.Atoi(qstId)
	if err != nil {
		return nil, err
	}

	repo.DB.Find(&result, id)

	if result.ID == 0 {
		return nil, nil
	}

	result.QuestionId = strconv.FormatUint(uint64(result.ID), 10)

	return &result, nil
}

func (repo *redisRepo) Create(ctx context.Context, qst model.Question) (*model.Question, error) {
	if !qst.IsQuestionValid() {
		return nil, fmt.Errorf("Question Can't Be Empty")
	}

	if !qst.IsUserIdValid() {
		return nil, fmt.Errorf("UserID Can't Be Empty")
	}

	elements := repo.DB.Keys(ctx, "QST*").Val()
	index := len(elements) + 1
	qst.QuestionId = fmt.Sprintf("QST%d", index)
	qst.ID = uint(index)

	err := repo.DB.Set(ctx, qst.QuestionId, qst, 0).Err()
	if err != nil {
		return nil, err
	}

	return repo.GetById(ctx, qst.QuestionId)
}

func (repo *mySqlRepo) Create(ctx context.Context, qst model.Question) (*model.Question, error) {
	result := repo.DB.Create(&qst)
	if result.Error != nil {
		return nil, result.Error
	}

	tmp, err := repo.GetById(ctx, strconv.Itoa(int(qst.ID)))
	if err != nil {
		return nil, err
	}

	return repo.Update(ctx, tmp.QuestionId, *tmp)
}

func (repo *redisRepo) Update(ctx context.Context, questionId string, qst model.Question) (*model.Question, error) {
	if !qst.IsQuestionValid() {
		return nil, fmt.Errorf("Question Can't Be Empty")
	}

	if !qst.IsUserIdValid() {
		return nil, fmt.Errorf("UserID Can't Be Empty")
	}

	tmp, err := repo.GetById(ctx, questionId)
	if err != nil {
		return nil, err
	} else if !tmp.IsIdValid() {
		return nil, fmt.Errorf("Question Not Found")
	} else if questionId != qst.QuestionId {
		return nil, fmt.Errorf("Question ID Does Not Match")
	}

	qst.ID = tmp.ID
	questionId = fmt.Sprintf("QST%s", questionId)
	err = repo.DB.Set(ctx, questionId, qst, 0).Err()
	if err != nil {
		return nil, err
	}

	return repo.GetById(ctx, qst.QuestionId)
}

func (repo *mySqlRepo) Update(ctx context.Context, questionId string, qst model.Question) (*model.Question, error) {
	tmp, err := repo.GetById(ctx, questionId)
	if err != nil {
		return nil, err
	}

	tmp.Question = qst.Question
	result := repo.DB.Save(&tmp)
	if result.Error != nil {
		return nil, result.Error
	}

	return repo.GetById(ctx, questionId)
}

func (repo *redisRepo) Delete(ctx context.Context, questionId string) (*model.Question, error) {
	result, err := repo.GetById(ctx, questionId)
	if err != nil {
		return nil, err
	}

	questionId = fmt.Sprintf("QST%s", questionId)
	_, err = repo.DB.Del(ctx, questionId).Result()
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (repo *mySqlRepo) Delete(ctx context.Context, questionId string) (*model.Question, error) {
	result, err := repo.GetById(ctx, questionId)
	if err != nil {
		return nil, err
	}

	id, err := strconv.Atoi(questionId)
	if err != nil {
		return nil, err
	}
	tmp := repo.DB.Delete(&model.Question{}, id)
	if tmp.Error != nil {
		return nil, tmp.Error
	}

	return result, nil
}

func (repo *redisRepo) Migrate(ctx context.Context) error {
	return nil
}

func (repo *mySqlRepo) Migrate(ctx context.Context) error {
	err := repo.DB.AutoMigrate(model.User{})
	if err != nil {
		return err
	}

	err = repo.DB.AutoMigrate(model.Question{})
	if err != nil {
		return err
	}

	err = repo.DB.AutoMigrate(model.Answer{})
	if err != nil {
		return err
	}

	return nil
}

func (repo *redisRepo) GetAnswerByQuestionId(ctx context.Context, questionId string) *model.Answer {
	for _, key := range repo.DB.Keys(ctx, "ANS*").Val() {
		tmp := model.Answer{}

		data, err := repo.DB.Get(ctx, key).Result()
		if err != nil {
			return nil
		}

		err = json.Unmarshal([]byte(data), &tmp)
		if err != nil {
			return nil
		} else {
			if tmp.QuestionId == questionId {
				return &tmp
			}
		}
	}

	return nil
}

func (repo *mySqlRepo) GetAnswerByQuestionId(ctx context.Context, questionId string) *model.Answer {
	result := []model.Answer{}
	repo.DB.Find(&result)

	for _, answer := range result {
		if answer.QuestionId == questionId {
			return &answer
		}
	}

	return nil
}

func (repo *redisRepo) GetQuestionsAnswersByUserId(ctx context.Context, userId string) *[]model.QuestionAnswer {
	questions := []model.QuestionAnswer{}
	for _, key := range repo.DB.Keys(ctx, "QST*").Val() {
		tmp := model.Question{}

		data, err := repo.DB.Get(ctx, key).Result()
		if err != nil {
			return nil
		}

		err = json.Unmarshal([]byte(data), &tmp)
		if err != nil {
			return nil
		} else {
			tmp.QuestionId = strings.Replace(tmp.QuestionId, "QST", "", -1)
			tmp.UserId = strings.Replace(tmp.UserId, "USR", "", -1)
			userId = strings.Replace(userId, "USR", "", -1)

			if tmp.UserId == userId {
				qa := model.QuestionAnswer{UserId: userId, QuestionId: tmp.QuestionId, Question: tmp.Question}
				an := repo.GetAnswerByQuestionId(ctx, tmp.QuestionId)

				if an != nil {
					an.AnswerId = strings.Replace(an.AnswerId, "ANS", "", -1)
					qa.Answer = *an
				}

				questions = append(questions, qa)
			}
		}
	}

	if len(questions) <= 0 {
		return nil
	}

	return &questions
}

func (repo *mySqlRepo) GetQuestionsAnswersByUserId(ctx context.Context, userId string) *[]model.QuestionAnswer {
	result := []model.QuestionAnswer{}

	questions := []model.Question{}
	repo.DB.Find(&questions)

	for _, question := range questions {
		if len(question.QuestionId) <= 0 {
			question.QuestionId = strconv.Itoa(int(question.ID))
		}
		if question.UserId == userId {
			qa := model.QuestionAnswer{UserId: userId, QuestionId: question.QuestionId, Question: question.Question}
			an := repo.GetAnswerByQuestionId(ctx, question.QuestionId)

			if an != nil {
				an.AnswerId = strings.Replace(an.AnswerId, "ANS", "", -1)
				qa.Answer = *an
			}

			result = append(result, qa)
		}
	}

	if len(result) <= 0 {
		return nil
	}

	return &result
}

func (repo *redisRepo) GetUserWithQuestionsAndAnswers(ctx context.Context, userId string) *model.UserQuestionsAnswers {
	if !strings.Contains(userId, "USR") {
		userId = "USR" + userId
	}

	data, err := repo.DB.Get(ctx, userId).Result()
	if err != nil {
		return nil
	}

	usr := model.User{}
	err = json.Unmarshal([]byte(data), &usr)
	if err != nil {
		return nil
	}

	usr.UserId = strings.Replace(usr.UserId, "USR", "", -1)
	result := model.UserQuestionsAnswers{UserId: usr.UserId, UserName: usr.UserName}
	qa := repo.GetQuestionsAnswersByUserId(ctx, usr.UserId)

	if qa != nil {
		result.QuestionsWithAnswers = *qa
	}

	return &result
}

func (repo *mySqlRepo) GetUserWithQuestionsAndAnswers(ctx context.Context, userId string) *model.UserQuestionsAnswers {
	usr := model.User{}
	id, err := strconv.Atoi(userId)
	if err != nil {
		return nil
	}

	repo.DB.Find(&usr, id)

	if usr.ID <= 0 {
		return nil
	}

	usr.UserId = strconv.FormatUint(uint64(usr.ID), 10)

	result := model.UserQuestionsAnswers{UserId: usr.UserId, UserName: usr.UserName}
	qa := repo.GetQuestionsAnswersByUserId(ctx, usr.UserId)

	if qa != nil {
		result.QuestionsWithAnswers = *qa
	}

	return &result
}
