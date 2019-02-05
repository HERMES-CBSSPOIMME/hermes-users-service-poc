package main

import (

	// Native Go Libs
	fmt "fmt"
	os "os"

	// 3rd Party Libs
	logruswrapper "github.com/terryvogelsang/logruswrapper"

	// Project Libs
	models "wave-demo-service-poc/models"
	router "wave-demo-service-poc/router"
)

var (

	// TODO: Change these to be fetched automatically with Kubernetes Secrets

	// MongoDBHost : MongoDB Host
	MongoDBHost = "wave-demo_mongodb"

	// MongoDBPort : MongoDB Port
	MongoDBPort = 27017

	// MongoDBUsername : MongoDB Username
	MongoDBUsername = "hermes-demo-user"

	// MongoDBPassword : MongoDB Password
	MongoDBPassword = "example"

	// MongoDBName : MongoDB Database Name
	MongoDBName = "hermesDemoDB"

	// MongoDBURL : MongoDB Connection URL
	MongoDBURL = fmt.Sprintf("mongodb://%s:%s@%s:%d/%s", MongoDBUsername, MongoDBPassword, MongoDBHost, MongoDBPort, MongoDBName)
)

func main() {

	// Get MongoDB communication interface
	// If an error occurs, program is set to panic
	mongoDB := models.NewMongoDB(MongoDBURL)

	env := &models.Env{
		MongoDB: mongoDB,
	}

	// Init logger to log in JSON Format in Stdout
	logruswrapper.Init(os.Stdout, "JSON")

	// Start Router
	router.Listen(env)
}
