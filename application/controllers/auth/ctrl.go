package controller

import (
	"encoding/json"
	"net/http"

	"be/application/controllers/auth/presenter"
	authFlow "be/usecase/auth"
)

func Login(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	var auth presenter.Auth
	json.NewDecoder(request.Body).Decode(&auth)
	result, err := authFlow.Login(auth.Username, auth.Password)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"error":"Internal server error"}`))
	}
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(result)
}

func Register(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	var auth presenter.Auth
	json.NewDecoder(request.Body).Decode(&auth)
	result, err := authFlow.Register(auth.Username, auth.Password)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"error":"Internal server error"}`))
	}
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(result)
}
