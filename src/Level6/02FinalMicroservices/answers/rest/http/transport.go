package http

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	mygrpc "BairesDev/Level6/02_Final_Microservices/answers/grpc/model"
	service "BairesDev/Level6/02_Final_Microservices/answers/rest/service"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/go-kit/log"
	"github.com/gorilla/mux"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewHttpServer(svc service.AnswerRESTService, logger log.Logger) *mux.Router {
	options := []httptransport.ServerOption{
		httptransport.ServerErrorLogger(logger),
		httptransport.ServerErrorEncoder(encodeErrorResponse),
	}

	//definition of a handlers
	getAllUsersHandler := httptransport.NewServer(
		MakeGetAll(svc),
		decodeGelAllRequest,
		encodeGetAllResponse,
		options...,
	)

	getUserByIdHandler := httptransport.NewServer(
		MakeGetById(svc),
		decodeGetByIdRequest,
		encodeGetByIdResponse,
		options...,
	)

	createUserHandler := httptransport.NewServer(
		MakeCreate(svc),
		decodeCreateRequest,
		encodeCreateResponse,
		options...,
	)

	updateUserHandler := httptransport.NewServer(
		MakeUpdate(svc),
		decodeUpdateRequest,
		encodeUpdateResponse,
		options...,
	)

	deleteUserHandler := httptransport.NewServer(
		MakeDelete(svc),
		decodeDeleteRequest,
		encodeDeleteResponse,
		options...,
	)

	r := mux.NewRouter()
	r.Methods("GET").Path("/api/Answer").Handler(getAllUsersHandler)
	r.Methods("GET").Path("/api/Answer/{AnswerID}").Handler(getUserByIdHandler)
	r.Methods("POST").Path("/api/Answer").Handler(createUserHandler)
	r.Methods("PUT").Path("/api/Answer/{AnswerID}").Handler(updateUserHandler)
	r.Methods("DELETE").Path("/api/Answer/{AnswerID}").Handler(deleteUserHandler)

	return r
}

func encodeErrorResponse(ctx context.Context, err error, w http.ResponseWriter) {
	if err == nil {
		panic("encodeError with nil error")
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}

func decodeGelAllRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var request getAllRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}

	return request, nil
}

func encodeGetAllResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

func decodeGetByIdRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var request getByIdRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}

	vars := mux.Vars(r)
	request.AnswerID = vars["AnswerID"]

	return request, nil
}

func encodeGetByIdResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

func decodeCreateRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var request createRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}

	usrId := r.Header.Get("UserID")
	if len(usrId) <= 0 {
		return nil, fmt.Errorf("Missing UserID Header")
	}

	qstId := r.Header.Get("QuestionID")
	if len(qstId) <= 0 {
		return nil, fmt.Errorf("Missing QuestionID Header")
	}

	if request.UserID != usrId {
		return nil, fmt.Errorf("UserId from Header Doesn't Match with UserId from Body")
	}

	if request.QuestionID != qstId {
		return nil, fmt.Errorf("QuestionId from Header Doesn't Match with QuestionId from Body")
	}

	conn, err := grpc.Dial("Answers:9003", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	defer conn.Close()
	c := mygrpc.NewAnswersClient(conn)

	usrqstans, err := c.GetUserWithQuestionsAndAnswers(ctx, &mygrpc.Input{UserId: usrId})
	if err != nil {
		return nil, err
	}

	if usrqstans == nil {
		return nil, fmt.Errorf("No User Found for UserID:%s", usrId)
	}

	for _, qa := range usrqstans.QuestionsWithAnswers {
		fmt.Println(fmt.Sprintf("Comparing %s and %s", qa.QuestionId, qstId))
		if qa.QuestionId == qstId {
			return request, nil
		}
	}

	return nil, fmt.Errorf("Question Not Found for UserId:%s", usrId)
}

func encodeCreateResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

func decodeUpdateRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var request updateRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}

	vars := mux.Vars(r)
	request.AnswerID = vars["AnswerID"]

	usrId := r.Header.Get("UserID")
	if len(usrId) <= 0 {
		return nil, fmt.Errorf("Missing UserID Header")
	}

	qstId := r.Header.Get("QuestionID")
	if len(qstId) <= 0 {
		return nil, fmt.Errorf("Missing QuestionID Header")
	}

	if request.UserID != usrId {
		return nil, fmt.Errorf("UserId from Header Doesn't Match with UserId from Body")
	}

	if request.QuestionID != qstId {
		return nil, fmt.Errorf("QuestionId from Header Doesn't Match with QuestionId from Body")
	}

	conn, err := grpc.Dial("Answers:9003", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	defer conn.Close()
	c := mygrpc.NewAnswersClient(conn)

	usrqstans, err := c.GetUserWithQuestionsAndAnswers(ctx, &mygrpc.Input{UserId: usrId})
	if err != nil {
		return nil, err
	}

	if usrqstans == nil {
		return nil, fmt.Errorf("No User Found for UserID:%s", usrId)
	}

	answer, err := c.GetAnswer(ctx, &mygrpc.Input{QuestionId: request.QuestionID})
	if err != nil {
		return nil, err
	}

	if request.UserID != answer.UserId || request.QuestionID != answer.QuestionId || request.AnswerID != answer.AnswerId {
		return nil, fmt.Errorf("UserID and/or QuestionID and/or AnswerID Doesn't Match")
	}

	for _, qa := range usrqstans.QuestionsWithAnswers {
		if qa.QuestionId == qstId {
			return request, nil
		}
	}

	return nil, fmt.Errorf("Question Not Found for UserId:%s", usrId)
}

func encodeUpdateResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

func decodeDeleteRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var request deleteRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}

	vars := mux.Vars(r)
	request.AnswerID = vars["AnswerID"]

	return request, nil
}

func encodeDeleteResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

type getAllRequest struct{}

type getByIdRequest struct {
	AnswerID string `json:"AnswerID"`
}

type createRequest struct {
	UserID     string `json:"UserID"`
	QuestionID string `json:"QuestionID"`
	Answer     string `json:"Answer"`
}

type updateRequest struct {
	AnswerID   string `json:"AnswerID"`
	QuestionID string `json:"QuestionID"`
	UserID     string `json:"UserID"`
	Answer     string `json:"Answer"`
}

type deleteRequest struct {
	AnswerID string `json:"AnswerID"`
}
