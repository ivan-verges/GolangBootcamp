package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/go-redis/redis/v9"
	"github.com/gorilla/mux"
)

type UserModel struct {
	UserId   string `json:"UserID"`
	UserName string `json:"UserName"`
}

func (user UserModel) IsIdValid() bool {
	return len(user.UserId) > 0 && strings.Contains(user.UserId, "USR")
}

func (user UserModel) IsUserValid() bool {
	return len(user.UserName) > 0
}

type QuestionModel struct {
	QuestionId string `json:"QuestionID"`
	UserId     string `json:"UserID"`
	Question   string `json:"Question"`
}

func (question QuestionModel) IsIdValid() bool {
	return len(question.QuestionId) > 0 && strings.Contains(question.QuestionId, "QST")
}

func (question QuestionModel) IsQuestionValid() bool {
	return len(question.Question) > 0 && strings.Contains(question.UserId, "USR")
}

type AnswerModel struct {
	AnswerId   string `json:"AnswerID"`
	UserId     string `json:"UserID"`
	QuestionId string `json:"QuestionID"`
	Answer     string `json:"Answer"`
}

func (answer AnswerModel) IsIdValid() bool {
	return len(answer.AnswerId) > 0 && strings.Contains(answer.AnswerId, "ANS")
}

func (answer AnswerModel) IsQuestionValid() bool {
	return len(answer.Answer) > 0 && strings.Contains(answer.UserId, "USR") && strings.Contains(answer.QuestionId, "QST")
}

type UserQuestionModel struct {
	User          UserModel       `json:"User"`
	UserQuestions []QuestionModel `json:"Questions"`
	UserAnswers   []AnswerModel   `json:"Answers"`
}

func (i QuestionModel) MarshalBinary() ([]byte, error) {
	return json.Marshal(i)
}

func (i AnswerModel) MarshalBinary() ([]byte, error) {
	return json.Marshal(i)
}

func (i UserModel) MarshalBinary() ([]byte, error) {
	return json.Marshal(i)
}

func (i UserQuestionModel) MarshalBinary() ([]byte, error) {
	return json.Marshal(i)
}

func GetUsers() []UserModel {
	result := []UserModel{}
	for _, key := range rdb.Keys(ctx, "USR*").Val() {
		tmp := UserModel{}

		data, err := rdb.Get(ctx, key).Result()
		if err != nil {
			panic(err)
		}

		err = json.Unmarshal([]byte(data), &tmp)
		if err != nil {
			panic(err)
		} else {
			result = append(result, tmp)
		}
	}

	return result
}

func GetUser(usrId string) UserModel {
	result := UserModel{}
	for _, usr := range GetUsers() {
		if usr.UserId == usrId {
			result = usr
		}
	}
	return result
}

func GetQuestions() []QuestionModel {
	result := []QuestionModel{}
	for _, key := range rdb.Keys(ctx, "QST*").Val() {
		tmp := QuestionModel{}

		data, err := rdb.Get(ctx, key).Result()
		if err != nil {
			panic(err)
		}

		err = json.Unmarshal([]byte(data), &tmp)
		if err != nil {
			panic(err)
		} else {
			result = append(result, tmp)
		}
	}

	return result
}

func GetQuestion(qstId string) QuestionModel {
	result := QuestionModel{}
	for _, qst := range GetQuestions() {
		if qst.QuestionId == qstId {
			result = qst
		}
	}
	return result
}

func GetAnswers() []AnswerModel {
	result := []AnswerModel{}
	for _, key := range rdb.Keys(ctx, "ANS*").Val() {
		tmp := AnswerModel{}

		data, err := rdb.Get(ctx, key).Result()
		if err != nil {
			panic(err)
		}

		err = json.Unmarshal([]byte(data), &tmp)
		if err != nil {
			panic(err)
		} else {
			result = append(result, tmp)
		}
	}

	return result
}

func GetAnswer(ansId string) AnswerModel {
	result := AnswerModel{}
	for _, ans := range GetAnswers() {
		if ans.AnswerId == ansId {
			result = ans
		}
	}
	return result
}

func GetUserQuestionAnswers() []UserQuestionModel {
	users := GetUsers()
	questions := GetQuestions()
	answers := GetAnswers()

	result := []UserQuestionModel{}
	for _, user := range users {
		tmp := UserQuestionModel{}
		tmp.User = user

		for _, question := range questions {
			if question.UserId == user.UserId {
				tmp.UserQuestions = append(tmp.UserQuestions, question)

				for _, answer := range answers {
					if answer.QuestionId == question.QuestionId {
						tmp.UserAnswers = append(tmp.UserAnswers, answer)
					}
				}
			}
		}

		result = append(result, tmp)
	}

	return result
}

func SaveUser(usr UserModel) error {
	if len(usr.UserName) <= 0 {
		return fmt.Errorf("UserName Can't Be Empty")
	}

	tmp := GetUser(usr.UserId)
	if len(tmp.UserId) <= 0 {
		usr.UserId = "USR" + strconv.FormatInt(time.Now().UnixMilli(), 10)
	}

	err := rdb.Set(ctx, usr.UserId, usr, 0).Err()
	if err != nil {
		return err
	}

	return nil
}

func SaveQuestion(usrId string, qst QuestionModel) error {
	if len(usrId) <= 0 {
		return fmt.Errorf("UserID Can't Be Empty")
	}

	if len(qst.Question) <= 0 {
		return fmt.Errorf("question can't be empty")
	}

	usr := GetUser(usrId)
	if len(usr.UserId) <= 0 {
		return fmt.Errorf("UserID {%s} Doesn't Exists", usrId)
	}

	tmp := GetQuestion(qst.QuestionId)
	if len(tmp.QuestionId) <= 0 {
		qst.QuestionId = "QST" + strconv.FormatInt(time.Now().UnixMilli(), 10)
	}

	qst.UserId = usrId
	err := rdb.Set(ctx, qst.QuestionId, qst, 0).Err()
	if err != nil {
		return err
	}

	return nil
}

func SaveAnswer(usrId string, qstId string, ans AnswerModel) error {
	if len(usrId) <= 0 {
		return fmt.Errorf("UserID Can't Be Empty")
	}

	if len(qstId) <= 0 {
		return fmt.Errorf("QuestionID Can't Be Empty")
	}

	if len(ans.Answer) <= 0 {
		return fmt.Errorf("answer can't be empty")
	}

	usr := GetUser(usrId)
	if len(usr.UserId) <= 0 {
		return fmt.Errorf("UserID {%s} Doesn't Exists", usrId)
	}

	qst := GetQuestion(qstId)
	if len(qst.QuestionId) <= 0 {
		return fmt.Errorf("QuestionID {%s} Doesn't Exists", qstId)
	}

	if qst.UserId != usr.UserId {
		return fmt.Errorf("user and question doesn't match")
	}

	tmp := GetAnswer(ans.AnswerId)
	if len(tmp.AnswerId) <= 0 {
		ans.AnswerId = "ANS" + strconv.FormatInt(time.Now().UnixMilli(), 10)
	}

	ans.QuestionId = qstId
	ans.UserId = usrId
	err := rdb.Set(ctx, ans.AnswerId, ans, 0).Err()
	if err != nil {
		return err
	}

	return nil
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	result := GetUsers()

	w.Header().Add("Content-Type", "application/json")
	if len(result) > 0 {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(result)
	} else {
		w.WriteHeader(http.StatusNoContent)
	}
}

func GetAllQuestions(w http.ResponseWriter, r *http.Request) {
	result := GetQuestions()

	w.Header().Add("Content-Type", "application/json")
	if len(result) > 0 {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(result)
	} else {
		w.WriteHeader(http.StatusNoContent)
	}
}

func GetAllAnswers(w http.ResponseWriter, r *http.Request) {
	result := GetAnswers()

	w.Header().Add("Content-Type", "application/json")
	if len(result) > 0 {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(result)
	} else {
		w.WriteHeader(http.StatusNoContent)
	}
}

func GetAllUserQuestionAnswer(w http.ResponseWriter, r *http.Request) {
	result := GetUserQuestionAnswers()

	w.Header().Add("Content-Type", "application/json")
	if len(result) > 0 {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(result)
	} else {
		w.WriteHeader(http.StatusNoContent)
	}
}

func GetUserById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["UserID"]

	result := UserModel{}
	for _, data := range GetUsers() {
		if data.UserId == id {
			result = data
		}
	}

	w.Header().Add("Content-Type", "application/json")
	if len(result.UserId) > 0 {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(result)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}

func GetQuestionById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["QuestionID"]

	result := QuestionModel{}
	for _, data := range GetQuestions() {
		if data.QuestionId == id {
			result = data
		}
	}

	w.Header().Add("Content-Type", "application/json")
	if len(result.QuestionId) > 0 {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(result)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}

func GetAnswerById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["AnswerID"]

	result := AnswerModel{}
	for _, data := range GetAnswers() {
		if data.AnswerId == id {
			result = data
		}
	}

	w.Header().Add("Content-Type", "application/json")
	if len(result.AnswerId) > 0 {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(result)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}

func GetUserQuestionAnswerById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["UserID"]

	result := UserQuestionModel{}
	for _, data := range GetUserQuestionAnswers() {
		if data.User.UserId == id {
			result = data
		}
	}

	w.Header().Add("Content-Type", "application/json")
	if len(result.User.UserId) > 0 {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(result)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	body, _ := io.ReadAll(r.Body)
	usr := UserModel{}
	json.Unmarshal([]byte(body), &usr)

	w.Header().Add("Content-Type", "application/json")
	err := SaveUser(usr)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(usr)
	}
}

func CreateQuestion(w http.ResponseWriter, r *http.Request) {
	usrId := r.Header.Get("UserID")

	if len(usrId) <= 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	body, _ := io.ReadAll(r.Body)
	qst := QuestionModel{}
	json.Unmarshal([]byte(body), &qst)

	w.Header().Add("Content-Type", "application/json")
	err := SaveQuestion(usrId, qst)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(qst)
	}
}

func CreateAnswer(w http.ResponseWriter, r *http.Request) {
	usrId := r.Header.Get("UserID")
	qstId := r.Header.Get("QuestionID")

	if len(usrId) <= 0 || len(qstId) <= 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	body, _ := io.ReadAll(r.Body)
	ans := AnswerModel{}
	json.Unmarshal([]byte(body), &ans)

	w.Header().Add("Content-Type", "application/json")
	err := SaveAnswer(usrId, qstId, ans)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(ans)
	}
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	usrId := vars["UserID"]
	if len(usrId) <= 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	body, _ := io.ReadAll(r.Body)
	usr := UserModel{}
	json.Unmarshal([]byte(body), &usr)

	tmp := GetUser(usrId)
	if len(tmp.UserId) <= 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	} else if tmp.UserId != usr.UserId {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	err := SaveUser(usr)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(usr)
	}
}

func UpdateQuestion(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	qstId := vars["QuestionID"]
	usrId := r.Header.Get("UserID")

	if len(usrId) <= 0 || len(qstId) <= 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	body, _ := io.ReadAll(r.Body)
	qst := QuestionModel{}
	json.Unmarshal([]byte(body), &qst)

	tmp := GetQuestion(qstId)
	if len(tmp.QuestionId) <= 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	} else if tmp.QuestionId != qst.QuestionId {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	err := SaveQuestion(usrId, qst)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(qst)
	}
}

func UpdateAnswer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ansId := vars["AnswerID"]
	usrId := r.Header.Get("UserID")
	qstId := r.Header.Get("QuestionID")

	if len(usrId) <= 0 || len(qstId) <= 0 || len(ansId) <= 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	body, _ := io.ReadAll(r.Body)
	ans := AnswerModel{}
	json.Unmarshal([]byte(body), &ans)

	tmp := GetAnswer(ansId)
	if len(tmp.AnswerId) <= 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	} else if tmp.AnswerId != ans.AnswerId {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	err := SaveAnswer(usrId, qstId, ans)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(ans)
	}
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["UserID"]

	_, err := rdb.Del(ctx, id).Result()
	if err != nil {
		panic(err)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func DeleteQuestion(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["QuestionID"]

	_, err := rdb.Del(ctx, id).Result()
	if err != nil {
		panic(err)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func DeleteAnswer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["AnswerID"]

	_, err := rdb.Del(ctx, id).Result()
	if err != nil {
		panic(err)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func CreateSampleData() {
	usrs := GetUsers()
	if len(usrs) == 0 {
		usrs = []UserModel{
			{UserName: "User 1"},
		}

		for _, usr := range usrs {
			SaveUser(usr)
		}
	}

	qsts := GetQuestions()
	if len(qsts) == 0 {
		usr := GetUsers()[0]
		qsts = []QuestionModel{
			{UserId: usr.UserId, Question: "Question 1"},
		}

		for _, qst := range qsts {
			SaveQuestion(qst.UserId, qst)
		}
	}

	anss := GetAnswers()
	if len(anss) == 0 {
		qst := GetQuestions()[0]
		anss = []AnswerModel{
			{UserId: qst.UserId, QuestionId: qst.QuestionId, Answer: "Answer 1"},
		}

		for _, ans := range anss {
			SaveAnswer(ans.UserId, ans.QuestionId, ans)
		}
	}
}

// TODO: Integrate GoKit
// TODO: Segment Project Into Micro-Services
// TODO: Move Context to Each Handler
var ctx context.Context
var rdb redis.UniversalClient

func main() {
	fmt.Println("Creatig Context and Redis DB Instances")
	ctx = context.Background()
	//rdb = redis.NewClient(&redis.Options{Addr: "localhost:6379", Password: "24071b57a13027c01339027dcccb98218f052a8f", DB: 0}) //Local
	rdb = redis.NewClient(&redis.Options{Addr: "redis:6379", Password: "24071b57a13027c01339027dcccb98218f052a8f", DB: 0}) //Docker

	fmt.Println("Creating Sample Data")
	CreateSampleData()

	fmt.Println("Serving on: http://localhost:8888")
	router := mux.NewRouter()

	//Get Methods
	router.HandleFunc("/User", GetAllUsers).Methods(http.MethodGet)
	router.HandleFunc("/User/{UserID}", GetUserById).Methods(http.MethodGet)
	router.HandleFunc("/Question", GetAllQuestions).Methods(http.MethodGet)
	router.HandleFunc("/Question/{QuestionID}", GetQuestionById).Methods(http.MethodGet)
	router.HandleFunc("/Answer", GetAllAnswers).Methods(http.MethodGet)
	router.HandleFunc("/Answer/{AnswerID}", GetAnswerById).Methods(http.MethodGet)
	router.HandleFunc("/AllInfo", GetAllUserQuestionAnswer).Methods(http.MethodGet)
	router.HandleFunc("/AllInfo/{UserID}", GetUserQuestionAnswerById).Methods(http.MethodGet)

	//Create Methods
	router.HandleFunc("/User", CreateUser).Methods(http.MethodPost)
	router.HandleFunc("/Question", CreateQuestion).Methods(http.MethodPost)
	router.HandleFunc("/Answer", CreateAnswer).Methods(http.MethodPost)

	//Update Methods
	router.HandleFunc("/User/{UserID}", UpdateUser).Methods(http.MethodPut)
	router.HandleFunc("/Question/{QuestionID}", UpdateQuestion).Methods(http.MethodPut)
	router.HandleFunc("/Answer/{AnswerID}", UpdateAnswer).Methods(http.MethodPut)

	//Delete Methods
	router.HandleFunc("/User/{UserID}", DeleteUser).Methods(http.MethodDelete)
	router.HandleFunc("/Question/{QuestionID}", DeleteQuestion).Methods(http.MethodDelete)
	router.HandleFunc("/Answer/{AnswerID}", DeleteAnswer).Methods(http.MethodDelete)

	log.Fatal(http.ListenAndServe(":8888", router))
}
