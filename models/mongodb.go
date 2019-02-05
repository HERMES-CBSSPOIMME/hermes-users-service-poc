package models

import (
	context "context"
	fmt "fmt"
	users "wave-demo-service-poc/users"
	utils "wave-demo-service-poc/utils"

	MongoBson "github.com/mongodb/mongo-go-driver/bson"
	mongo "github.com/mongodb/mongo-go-driver/mongo"
	bson "gopkg.in/mgo.v2/bson"
)

const (

	// HermesDatabaseName : Database name of the Hermes project as defined in MongoDB
	HermesDatabaseName = "hermesDemoDB"

	// UserCollection : MongoDB Collection containing users profile
	UserCollection = "user-collection"
)

// MongoDBInterface : MongoDB Communication interface
type MongoDBInterface interface {
	// User related functions
	AddUser(user *users.User) error
	GetUserById(uid string) (*users.User, error)
	GetUserByUsername(username string) (*users.User, error)
	UpdateUser(user *users.User) error
	DeleteUser(uid string) error
}

// MongoDB : MongoDB communication interface
type MongoDB struct {
	Client         *mongo.Client
	HermesDB       *mongo.Database
	UserCollection *mongo.Collection
}

// NewMongoDB : Return a new MongoDB abstraction struct
func NewMongoDB(connectionURL string) *MongoDB {

	// Get connection to DB
	client, err := mongo.NewClient(connectionURL)

	if err != nil {
		utils.PanicOnError(err, "Failed to connect to MongoDB")
	}

	err = client.Connect(context.TODO())

	if err != nil {
		utils.PanicOnError(err, "Failed to connect to context")
	}

	// Get database reference
	hermesDB := client.Database(HermesDatabaseName)

	// Get collections references
	userCollection := hermesDB.Collection(UserCollection)

	// Return new MongoDB abstraction struct
	return &MongoDB{
		Client:         client,
		HermesDB:       hermesDB,
		UserCollection: userCollection,
	}
}

// AddUser: Add user entry in database
func (mongoDB *MongoDB) AddUser(user *users.User) error {

	// Marshal struct into bson object
	doc, err := bson.Marshal(*user)

	if err != nil {
		return err
	}

	// Insert group conversation into DB
	_, err = mongoDB.UserCollection.InsertOne(nil, doc)

	if err != nil {
		return err
	}

	return nil
}

func (mongoDB *MongoDB) UpdateUser(user *users.User) error {

	update := MongoBson.NewDocument(
		MongoBson.EC.SubDocumentFromElements(
			"$set",
			MongoBson.EC.String("username", user.Username),
			MongoBson.EC.String("name", user.Name),
			MongoBson.EC.String("surname", user.Surname),
			MongoBson.EC.String("email", user.Email),
			MongoBson.EC.String("pictureURL", user.Picture_URL),
			MongoBson.EC.String("password", user.Password),
		),
	)

	_, err := mongoDB.UserCollection.UpdateOne(nil, bson.M{"_id": user.Uid}, update)

	if err != nil {
		fmt.Println("ERR 2")
		return err
	}

	return nil
}

func (mongoDB *MongoDB) GetUserById(uid string) (*users.User, error) {
	var dr *mongo.DocumentResult
	var user users.User

	dr = mongoDB.UserCollection.FindOne(nil, bson.M{"_id": uid})

	err := dr.Decode(&user)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (mongoDB *MongoDB) DeleteUser(uid string) error {

	_, err := mongoDB.UserCollection.DeleteOne(nil, bson.M{"_id": uid})

	if err != nil {
		return err
	}

	return nil
}

func (mongoDB *MongoDB) GetUserByUsername(username string) (*users.User, error) {
	var dr *mongo.DocumentResult
	var user users.User

	dr = mongoDB.UserCollection.FindOne(nil, bson.M{"username": username})

	err := dr.Decode(&user)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

/*
func (mongoDB *MongoDB) GetAllUsers() ([]users.User, error) {
	var users []users.User

	cur, err := mongoDB.UserCollection.Find(nil, bson.M{})

	for cur.Next(nil) {
		elem := MongoBson.NewDocument()
		if err := cur.Decode(elem); err != nil {
			log.Fatal(err)
		}
	}
}*/
