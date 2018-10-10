package router

import (

	// Native Go Libs
	http "net/http"

	// 3rd Party Libs
	gocustomhttpresponse "github.com/terryvogelsang/gocustomhttpresponse"
	logruswrapper "github.com/terryvogelsang/logruswrapper"
)

type Greeter struct {
	Message string
}

// HelloWorld : A Simple HelloWorld Endpoint
func HelloWorld(w http.ResponseWriter, r *http.Request) error {

	// Logging demo
	log := logruswrapper.NewEntry("UsersService", "/helloworld", logruswrapper.CodeSuccess)
	logruswrapper.Info(log)

	greeter := Greeter{Message: "Hello World"}

	gocustomhttpresponse.WriteResponse(greeter, log, w)
	return nil
}

func GetUser(w http.ResponseWriter, r *http.Request) error {

}

func GetAllUsers(w http.ResponseWriter, r *http.Request) error {

}
