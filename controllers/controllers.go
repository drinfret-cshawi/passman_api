package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type JsonResponse struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}

func handleError(err error, w http.ResponseWriter) {
	if err != nil {
		fmt.Println(err)
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		var response = JsonResponse{Type: "error", Message: fmt.Sprint(err)}
		_ = json.NewEncoder(w).Encode(response)
	}
}

func handleJSONResponse[TResponse any](response TResponse, w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Headers, Origin,Accept, X-Requested-With, Access-Control-Request-Method, Access-Control-Request-Headers")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		panic(err)
	}
}

func GetOne(w http.ResponseWriter, _ *http.Request) {
	var response = JsonResponse{Type: "success", Message: "One"}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		panic(err)
	}

}

func GetTwo(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("<html><body><h1>Allo!</h1></body></html>"))
	if err != nil {
		return
	}

}
