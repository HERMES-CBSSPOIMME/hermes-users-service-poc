package router

import (

	// Native Go Libs
	fmt "fmt"
	http "net/http"

	// Project Libs
	models "hermes-users-service/models"
	handlers "hermes-users-service/router/handlers"

	customHandle "github.com/HERMES-CBSSPOIMME/hermes-go-utils-lib/customhandle"

	// 3rd Party Libs

	mux "github.com/gorilla/mux"
	cors "github.com/rs/cors"
)

const (
	// PORT : Listening Port
	PORT int = 8086
)

// Listen : Defines all router routing rules and handlers.
// Serves the API at defined port constant.
func Listen(env *models.Env) {

	r := mux.NewRouter().StrictSlash(false)

	v1 := r.PathPrefix("/v1").Subrouter()
	// HelloWorld Endpoint
	helloWorldV1 := v1.PathPrefix("/helloworld").Subrouter()
	helloWorldV1.Handle("", customHandle.CustomHandle(handlers.HelloWorld)).Methods("GET")

	usersV1 := v1.PathPrefix("/users").Subrouter()
	usersV1.Handle("", handlers.CustomHandle(env, handlers.CreateNewUser)).Methods("POST")
	usersV1.Handle("/{uid}", handlers.CustomHandle(env, handlers.GetUser)).Methods("GET")
	usersV1.Handle("/{uid}", handlers.CustomHandle(env, handlers.UpdateUser)).Methods("PUT")
	usersV1.Handle("/{uid}", handlers.CustomHandle(env, handlers.DeleteUser)).Methods("DELETE")
	usersV1.Handle("/login", handlers.CustomHandle(env, handlers.Login)).Methods("POST")

	corsHandler := cors.New(cors.Options{
		AllowedHeaders:   []string{"X-Requested-With"},
		AllowedOrigins:   []string{"http://frontend.localhost"},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "HEAD", "POST", "PUT", "OPTIONS"},
	})

	http.ListenAndServe(":"+fmt.Sprintf("%d", PORT), corsHandler.Handler(r))
}
