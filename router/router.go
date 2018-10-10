package router

import (

	// Native Go Libs
	fmt "fmt"
	http "net/http"

	// Project Libs
	handlers "hermes-users-service/router/handlers"

	customHandle "github.com/HERMES-CBSSPOIMME/hermes-go-utils-lib/customhandle"

	// 3rd Party Libs

	mux "github.com/gorilla/mux"
	cors "github.com/rs/cors"
)

const (
	// PORT : Listening Port
	PORT int = 8085
)

// Listen : Defines all router routing rules and handlers.
// Serves the API at defined port constant.
func Listen() {

	r := mux.NewRouter().StrictSlash(false)

	v1 := r.PathPrefix("/v1").Subrouter()

	// HelloWorld Endpoint
	helloWorldV1 := v1.PathPrefix("/helloworld").Subrouter()
	helloWorldV1.Handle("", customHandle.CustomHandle(handlers.HelloWorld)).Methods("GET")

	corsHandler := cors.New(cors.Options{
		AllowedHeaders:   []string{"X-Requested-With"},
		AllowedOrigins:   []string{"http://frontend.localhost"},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "HEAD", "POST", "PUT", "OPTIONS"},
	})

	http.ListenAndServe(":"+fmt.Sprintf("%d", PORT), corsHandler.Handler(r))
}
