package model

import (
	"encoding/json"
	"fmt"
)

type User struct {
	ID       uint   `gorm:"primaryKey"`
	UserId   string `gorm:"-:all"`
	UserName string `json:"UserName"`
}

func (user User) IsIdValid() bool {
	return len(user.UserId) > 0
}

func (user User) IsUserValid() bool {
	return len(user.UserName) > 0
}

func (user User) MarshalBinary() ([]byte, error) {
	return json.Marshal(user)
}

type Question struct {
	ID         uint   `gorm:"primaryKey"`
	QuestionId string `gorm:"-:all"`
	UserId     string `json:"UserID"`
	Question   string `json:"Question"`
}

func (question Question) IsIdValid() bool {
	return len(question.QuestionId) > 0
}

func (question Question) IsQuestionValid() bool {
	return len(question.Question) > 0
}

func (question Question) IsUserIdValid() bool {
	return len(question.UserId) > 0
}

func (question Question) MarshalBinary() ([]byte, error) {
	return json.Marshal(question)
}

type Answer struct {
	ID         uint   `gorm:"primaryKey"`
	AnswerId   string `gorm:"-:all"`
	QuestionId string `json:"QuestionID"`
	UserId     string `json:"UserID"`
	Answer     string `json:"Answer"`
}

func (answer Answer) IsIdValid() bool {
	return len(answer.AnswerId) > 0
}

func (answer Answer) IsAnswerValid() bool {
	return len(answer.Answer) > 0
}

func (answer Answer) IsUserIdValid() bool {
	return len(answer.UserId) > 0
}

func (answer Answer) IsQuestionIdValid() bool {
	return len(answer.QuestionId) > 0
}

func (answer Answer) MarshalBinary() ([]byte, error) {
	return json.Marshal(answer)
}

type QuestionAnswer struct {
	QuestionId string `json:"QuestionID"`
	UserId     string `json:"UserID"`
	Question   string `json:"Question"`
	Answer     Answer `json:"Answer"`
}

type UserQuestionsAnswers struct {
	UserId               string           `json:"UserID"`
	UserName             string           `json:"UserName"`
	QuestionsWithAnswers []QuestionAnswer `json:"QuestionsWithAnswers"`
}

type SqlConnectionString struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

func (connection SqlConnectionString) GetShortConnectionString() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", connection.User, connection.Password, connection.Host, connection.Port, connection.Database)
}

func GetMySqlConnection() SqlConnectionString {
	return SqlConnectionString{Host: "mysql", Port: "3306", User: "mysqluser", Password: "MyP4ssw0rd", Database: "microservices"} //Docker
	//return SqlConnectionString{Host: "localhost", Port: "3306", User: "mysqluser", Password: "MyP4ssw0rd", Database: "microservices"} //Local
}

type RedisConnectionString struct {
	Host     string
	Port     string
	Password string
	Instance int
}

func (connection RedisConnectionString) GetAddress() string {
	return fmt.Sprintf("%s:%s", connection.Host, connection.Port)
}

func GetRedisConnection() RedisConnectionString {
	return RedisConnectionString{Host: "redis", Port: "6379", Password: "24071b57a13027c01339027dcccb98218f052a8f", Instance: 0} //Docker
	//return RedisConnectionString{Host: "localhost", Port: "6379", Password: "24071b57a13027c01339027dcccb98218f052a8f", Instance: 0} //Local
}
