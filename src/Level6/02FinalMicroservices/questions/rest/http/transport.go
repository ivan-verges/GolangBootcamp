package http

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	mygrpc "BairesDev/Level6/02_Final_Microservices/questions/grpc/model"
	service "BairesDev/Level6/02_Final_Microservices/questions/rest/service"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/go-kit/log"
	"github.com/gorilla/mux"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewHttpServer(svc service.QuestionRESTService, logger log.Logger) *mux.Router {
	options := []httptransport.ServerOption{
		httptransport.ServerErrorLogger(logger),
		httptransport.ServerErrorEncoder(encodeErrorResponse),
	}

	//definition of a handlers
	getAllHandler := httptransport.NewServer(
		MakeGetAll(svc),
		decodeGelAllRequest,
		encodeGetAllResponse,
		options...,
	)

	getByIdHandler := httptransport.NewServer(
		MakeGetById(svc),
		decodeGetByIdRequest,
		encodeGetByIdResponse,
		options...,
	)

	createHandler := httptransport.NewServer(
		MakeCreate(svc),
		decodeCreateRequest,
		encodeCreateResponse,
		options...,
	)

	updateHandler := httptransport.NewServer(
		MakeUpdate(svc),
		decodeUpdateRequest,
		encodeUpdateResponse,
		options...,
	)

	deleteHandler := httptransport.NewServer(
		MakeDelete(svc),
		decodeDeleteRequest,
		encodeDeleteResponse,
		options...,
	)

	r := mux.NewRouter()
	r.Methods("GET").Path("/api/Question").Handler(getAllHandler)
	r.Methods("GET").Path("/api/Question/{QuestionID}").Handler(getByIdHandler)
	r.Methods("POST").Path("/api/Question").Handler(createHandler)
	r.Methods("PUT").Path("/api/Question/{QuestionID}").Handler(updateHandler)
	r.Methods("DELETE").Path("/api/Question/{QuestionID}").Handler(deleteHandler)

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
	request.QuestionID = vars["QuestionID"]

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

	if request.UserID != usrId {
		return nil, fmt.Errorf("UserId from Header Doesn't Match with UserId from Body")
	}

	conn, err := grpc.Dial("Questions:9002", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	defer conn.Close()
	c := mygrpc.NewQuestionsClient(conn)

	usrqstans, err := c.GetUserWithQuestionsAndAnswers(ctx, &mygrpc.Input{UserId: usrId})
	if err != nil {
		return nil, err
	}

	if usrqstans == nil {
		return nil, fmt.Errorf("No User Found for UserID:%s", usrId)
	}

	return request, nil
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
	request.QuestionID = vars["QuestionID"]

	usrId := r.Header.Get("UserID")
	if len(usrId) <= 0 {
		return nil, fmt.Errorf("Missing UserID Header")
	}

	if request.UserID != usrId {
		return nil, fmt.Errorf("UserId from Header Doesn't Match with UserId from Body")
	}

	conn, err := grpc.Dial("Questions:9002", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	defer conn.Close()
	c := mygrpc.NewQuestionsClient(conn)

	usrqstans, err := c.GetUserWithQuestionsAndAnswers(ctx, &mygrpc.Input{UserId: usrId})
	if err != nil {
		return nil, err
	}

	if usrqstans == nil {
		return nil, fmt.Errorf("No User Found for UserID:%s", usrId)
	}

	question, err := c.GetQuestion(ctx, &mygrpc.Input{QuestionId: request.QuestionID})
	if err != nil {
		return nil, err
	}

	if request.UserID != question.UserId || request.QuestionID != question.QuestionId {
		return nil, fmt.Errorf("UserID and/or Question ID Doesn't Match")
	}

	return request, nil
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
	request.QuestionID = vars["QuestionID"]

	return request, nil
}

func encodeDeleteResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

type getAllRequest struct{}

type getByIdRequest struct {
	QuestionID string `json:"QuestionID"`
}

type createRequest struct {
	UserID   string `json:"UserID"`
	Question string `json:"Question"`
}

type updateRequest struct {
	QuestionID string `json:"QuestionID"`
	UserID     string `json:"UserID"`
	Question   string `json:"Question"`
}

type deleteRequest struct {
	QuestionID string `json:"QuestionID"`
}
