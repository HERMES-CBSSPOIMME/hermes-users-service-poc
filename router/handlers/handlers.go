package router

import (

	// Native Go Libs
	json "encoding/json"
	ioutil "io/ioutil"
	http "net/http"

	//strings "strings"

	// 3rd Party Libs
	mux "github.com/gorilla/mux"
	gocustomhttpresponse "github.com/terryvogelsang/gocustomhttpresponse"
	logruswrapper "github.com/terryvogelsang/logruswrapper"

	// project intern includes
	auth "wave-demo-service-poc/auth"
	models "wave-demo-service-poc/models"
	users "wave-demo-service-poc/users"
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

// GetUser :
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

func AuthenticateUser(env *models.Env, w http.ResponseWriter, r *http.Request) error {
	token := r.Header.Get("token")
	//splitToken := strings.Split(token, "Bearer ")
	//token = splitToken[1]

	uw, err := auth.ValidateToken(token)
	if err != nil {
		errorLog := logruswrapper.NewEntry("UsersService", "Auth", logruswrapper.CodeInvalidToken)
		gocustomhttpresponse.WriteResponse(nil, errorLog, w)
		return err
	}

	//log := logruswrapper.NewEntry("UsersService", "Auth", logruswrapper.CodeSuccess)

	//gocustomhttpresponse.WriteResponse(uw, log, w)
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(&uw)

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

	log := logruswrapper.NewEntry("UsersService", "AddUser", logruswrapper.CodeSuccess)

	gocustomhttpresponse.WriteResponse(user.Uid, log, w)

	return nil
}

func DeleteUser(env *models.Env, w http.ResponseWriter, r *http.Request) error {
	vars := mux.Vars(r)
	uid := vars["uid"]

	err := env.MongoDB.DeleteUser(uid)

	if err != nil {
		errorLog := logruswrapper.NewEntry("UsersService", "DeleteUser", logruswrapper.CodeInvalidJSON)
		gocustomhttpresponse.WriteResponse(nil, errorLog, w)
		return err
	}

	log := logruswrapper.NewEntry("UsersService", "DeleteUser", logruswrapper.CodeSuccess)

	gocustomhttpresponse.WriteResponse(uid, log, w)

	return nil
}

func UpdateUser(env *models.Env, w http.ResponseWriter, r *http.Request) error {
	temp, _ := ioutil.ReadAll(r.Body)
	vars := mux.Vars(r)
	uid := vars["uid"]

	var user users.User

	err := json.Unmarshal(temp, &user)
	if err != nil {
		errorLog := logruswrapper.NewEntry("UsersService", "UpdateUser", logruswrapper.CodeInvalidJSON)
		gocustomhttpresponse.WriteResponse(nil, errorLog, w)
		return err
	}

	user.Uid = uid

	err = env.MongoDB.UpdateUser(&user)
	if err != nil {
		errorLog := logruswrapper.NewEntry("UsersService", "UpdateUser", logruswrapper.CodeInvalidJSON)
		gocustomhttpresponse.WriteResponse(nil, errorLog, w)
		return err
	}

	log := logruswrapper.NewEntry("UsersService", "UdateUser", logruswrapper.CodeSuccess)

	gocustomhttpresponse.WriteResponse(user.Uid, log, w)

	return nil
}

func Login(env *models.Env, w http.ResponseWriter, r *http.Request) error {
	temp, _ := ioutil.ReadAll(r.Body)

	var creds auth.Credentials

	err := json.Unmarshal(temp, &creds)
	if err != nil {
		errorLog := logruswrapper.NewEntry("UsersService", "Login", logruswrapper.CodeBadLogin)
		gocustomhttpresponse.WriteResponse(nil, errorLog, w)
		return err
	}

	uid, err := creds.Verify(env)
	if err != nil || uid == "" {
		errorLog := logruswrapper.NewEntry("UsersService", "Login", logruswrapper.CodeBadLogin)
		gocustomhttpresponse.WriteResponse(nil, errorLog, w)
		return err
	}

	tw, err := auth.CreateToken(uid)
	if err != nil {
		errorLog := logruswrapper.NewEntry("UsersService", "Login", logruswrapper.CodeBadLogin)
		gocustomhttpresponse.WriteResponse(nil, errorLog, w)
		return err
	}

	log := logruswrapper.NewEntry("UsersService", "Login", logruswrapper.CodeSuccess)

	gocustomhttpresponse.WriteResponse(tw, log, w)

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
