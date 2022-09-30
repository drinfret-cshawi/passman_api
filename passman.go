package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"log"
	"net/http"
	"passman_api/controllers"
)

const port = 8000

func main() {
	fmt.Println("Setting up routes and starting server...")

	router := mux.NewRouter()
	// v1 test routes
	router.HandleFunc("/v1/one", controllers.GetOne).Methods("GET")
	router.HandleFunc("/v1/two", controllers.GetTwo).Methods("GET")

	// v1 users routes
	router.HandleFunc("/v1/users/", controllers.GetUsers).Methods("GET")
	router.HandleFunc("/v1/users/{id:[0-9]+}", controllers.GetUserById).Methods("GET")
	router.HandleFunc("/v1/users/", controllers.AddUser).Methods("POST")
	router.HandleFunc("/v1/users/{id:[0-9]+}", controllers.DeleteUser).Methods("DELETE")
	router.HandleFunc("/v1/users/{id:[0-9]+}", controllers.UpdateUser).Methods("PUT")

	// v1 passwords routes
	router.HandleFunc("/v1/passwords/{user_id:[0-9]+}", controllers.GetPasswordsForUser).Methods("GET")
	router.HandleFunc("/v1/passwords/{user_id:[0-9]+}", controllers.AddPassword).Methods("POST")
	router.HandleFunc("/v1/passwords/{user_id:[0-9]+}", controllers.DeletePassword).Methods("DELETE")
	router.HandleFunc("/v1/passwords/{user_id:[0-9]+}", controllers.UpdatePassword).Methods("PUT")

	fmt.Printf("Server listening at :%d\n", port)
	c := cors.New(cors.Options{
		//AllowedOrigins:   []string{"http://foo.com", "http://foo.com:8080"},
		AllowedOrigins:   []string{"http://localhost*", "https://localhost*", "http://127.0.0.1:53322"},
		AllowCredentials: true,
		Debug:            true,
	})
	//handler := cors.Default().Handler(router)
	handler := c.Handler(router)
	//handler = http.TimeoutHandler(handler, time.Second, "Timeout!")

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), handler))
}
