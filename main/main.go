package main

import (

	// Native Go Libs
	os "os"

	// 3rd Party Libs
	logruswrapper "github.com/terryvogelsang/logruswrapper"

	// Project Libs
	router "hermes-users-service/router"
)

func main() {

	// Init logger to log in JSON Format in Stdout
	logruswrapper.Init(os.Stdout, "JSON")

	// Start Router
	router.Listen()
}
