package models

import (
	context "context"
	users "hermes-users-service/users"
	utils "hermes-users-service/utils"

	//mongoBSON "github.com/mongodb/mongo-go-driver/bson"
	mongo "github.com/mongodb/mongo-go-driver/mongo"
	bson "gopkg.in/mgo.v2/bson"
)

const (

	// HermesDatabaseName : Database name of the Hermes project as defined in MongoDB
	HermesDatabaseName = "hermesDemoDB"

	// UserCollection : MongoDB Collection containing ussers profile
	UserCollection = "user-collection"
)

// MongoDBInterface : MongoDB Communication interface
type MongoDBInterface interface {
	AddUser(user *users.User) error
	GetUserById(uid string) (*users.User, error)
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

func (mongoDB *MongoDB) GetAllUsers() ([]users.User, error) {
	var users []users.User

	cur, err := mongoDB.UserCollection.Find(nil, bson.M{})

	for cur.Next(nil) {
		elem := bson.NewDocument()
		if err := cur.Decode(elem); err != nil {
			log.Fatal(err)
		}
	}
}
