package http

import (
	"context"
	"encoding/json"
	"net/http"

	service "BairesDev/Level6/02_Final_Microservices/users/rest/service"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/go-kit/log"
	"github.com/gorilla/mux"
)

func NewHttpServer(svc service.UserRESTService, logger log.Logger) *mux.Router {
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
	r.Methods("GET").Path("/api/User").Handler(getAllUsersHandler)
	r.Methods("GET").Path("/api/User/{UserID}").Handler(getUserByIdHandler)
	r.Methods("POST").Path("/api/User").Handler(createUserHandler)
	r.Methods("PUT").Path("/api/User/{UserID}").Handler(updateUserHandler)
	r.Methods("DELETE").Path("/api/User/{UserID}").Handler(deleteUserHandler)

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
	request.UserID = vars["UserID"]

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
	request.UserID = vars["UserID"]

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
	request.UserID = vars["UserID"]

	return request, nil
}

func encodeDeleteResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

type getAllRequest struct{}

type getByIdRequest struct {
	UserID string `json:"UserID"`
}

type createRequest struct {
	UserName string `json:"UserName"`
}

type updateRequest struct {
	UserID   string `json:"UserID"`
	UserName string `json:"UserName"`
}

type deleteRequest struct {
	UserID string `json:"UserID"`
}
