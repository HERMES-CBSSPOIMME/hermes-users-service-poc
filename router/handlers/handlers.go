package router

import (

	// Native Go Libs
	json "encoding/json"
	ioutil "io/ioutil"
	http "net/http"

	// 3rd Party Libs
	mux "github.com/gorilla/mux"
	gocustomhttpresponse "github.com/terryvogelsang/gocustomhttpresponse"
	logruswrapper "github.com/terryvogelsang/logruswrapper"

	models "hermes-users-service/models"
	users "hermes-users-service/users"
)

type Greeter struct {
	Message string
}

type (
	// Handler : Custom type to work with CustomHandle wrapper
	Handler func(env *models.Env, w http.ResponseWriter, r *http.Request) error
)

// HelloWorld : A Simple HelloWorld Endpoint
func HelloWorld(w http.ResponseWriter, r *http.Request) error {

	// Logging demo
	log := logruswrapper.NewEntry("UsersService", "/helloworld", logruswrapper.CodeSuccess)
	logruswrapper.Info(log)

	greeter := Greeter{Message: "Hello World"}

	gocustomhttpresponse.WriteResponse(greeter, log, w)
	return nil
}

func GetUser(env *models.Env, w http.ResponseWriter, r *http.Request) error {
	vars := mux.Vars(r)
	uid := vars["uid"]

	user, err := env.MongoDB.GetUserById(uid)

	if err != nil {
		errorLog := logruswrapper.NewEntry("UsersService", "GetUser", logruswrapper.CodeInvalidJSON)
		gocustomhttpresponse.WriteResponse(nil, errorLog, w)
		return err
	}

	json.NewEncoder(w).Encode(user)

	return nil
}

func GetAllUsers(env *models.Env, w http.ResponseWriter, r *http.Request) error {

	return nil
}

func AuthenticateUser(env *models.Env, w http.ResponseWriter, r *http.Request) error {
	return nil
}

func CreateNewUser(env *models.Env, w http.ResponseWriter, r *http.Request) error {

	temp, _ := ioutil.ReadAll(r.Body)

	var user users.User

	err := json.Unmarshal(temp, &user)
	if err != nil {
		errorLog := logruswrapper.NewEntry("UsersService", "CreateNewUser", logruswrapper.CodeInvalidJSON)
		gocustomhttpresponse.WriteResponse(nil, errorLog, w)
		return err
	}

	users.AssignId(&user)

	err = env.MongoDB.AddUser(&user)

	if err != nil {
		errorLog := logruswrapper.NewEntry("UsersService", "CreateNewUser", logruswrapper.CodeInvalidJSON)
		gocustomhttpresponse.WriteResponse(nil, errorLog, w)
		return err
	}

	log := logruswrapper.NewEntry("UsersService", "/helloworld", logruswrapper.CodeSuccess)

	gocustomhttpresponse.WriteResponse(user.Uid, log, w)

	return nil
}

func DeleteUser(env *models.Env, w http.ResponseWriter, r *http.Request) error {
	return nil
}

func UpdateUser(env *models.Env, w http.ResponseWriter, r *http.Request) error {
	return nil
}

func Login(env *models.Env, w http.ResponseWriter, r *http.Request) error {
	return nil
}

func Logout(env *models.Env, w http.ResponseWriter, r *http.Request) error {
	return nil
}

// CustomHandle : Custom Handlers Wrapper for API
func CustomHandle(env *models.Env, handlers ...Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		for _, h := range handlers {
			err := h(env, w, r)
			if err != nil {
				errorLog := logruswrapper.NewEntry("UsersService", "/something", err.Error())
				gocustomhttpresponse.WriteResponse(nil, errorLog, w)
				return
			}
		}
	})
}
