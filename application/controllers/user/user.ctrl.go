package controller

import (
	"encoding/json"
	"net/http"

	userFlow "be/usecase/user"
)

func GetAll(response http.ResponseWriter, request *http.Request) {

	response.Header().Set("Content-Type", "application/json")
	users, err := userFlow.List()
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"error":"Internal server error"}`))
	}
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(users)
}
