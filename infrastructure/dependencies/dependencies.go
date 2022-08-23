// Package dependencies provider dependencies
package dependencies

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/apex/log"
	"github.com/apex/log/handlers/cli"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"lolesports/infrastructure/configuration"
	"lolesports/utils"
)

// Container container for dependencies injection in modules
type Container struct {
	logger *utils.Handler

	clock          Clock
	mongoDatastore MongoDatastore
}

type (
	// Clock defines a clock functionality
	Clock interface {
		Now() time.Time
	}

	clock struct {
		Clock
	}

	MongoDatastore interface {
		Connect(config configuration.Configuration) error
		Disconnect()
		DB() *mongo.Database
		Client() *mongo.Client
	}

	mongoDatastore struct {
		db     mongo.Database
		client mongo.Client
	}
)

// StartDependencies init function for dependencies
// It returns a dependencies container with all the dependencies started
func StartDependencies() *Container {
	// Clock
	clock := NewClock()

	// Logger
	logger := &utils.Handler{
		Handlers: make([]log.Handler, 0),
	}
	logger = utils.New(cli.New(os.Stderr))

	// db
	mongoDS := NewMongoDataStore()

	container := &Container{
		clock:          clock,
		logger:         logger,
		mongoDatastore: mongoDS,
	}

	// done
	return container
}

// Logger returns logging service for App
func (container *Container) Logger() utils.Handler {
	return *container.logger
}

// Clock return clock service
func (container *Container) Clock() Clock {
	return container.clock
}

// MongoDataStore return MongoDataStore
func (container *Container) MongoDataStore() MongoDatastore {
	return container.mongoDatastore
}

// MongoDatabase return MongoDatabase
func (container *Container) MongoDataBase() *mongo.Database {
	return container.mongoDatastore.DB()
}

// MongoClient return Client
func (container *Container) MongoClient() *mongo.Client {
	return container.mongoDatastore.Client()
}

// NewClock returns a new instance of clock object
func NewClock() Clock {
	return &clock{}
}

// Now Returns a timestamp
func (clock *clock) Now() time.Time {
	return time.Now().In(time.FixedZone("UTC", 0))
}

// NewMongoDataStore returns a new instance of mongoDatastore object
func NewMongoDataStore() MongoDatastore {
	return &mongoDatastore{}
}

func (mongoDatastore *mongoDatastore) Connect(config configuration.Configuration) error {
	url := fmt.Sprintf("mongodb://%s:%s@%s:%s/?authSource=admin&readPreference=primary&directConnection=true&ssl=false", config.DBUserName(), config.DBPassword(), config.DBHost(), config.DBPort())
	client, err := mongo.NewClient(options.Client().ApplyURI(url))
	if err != nil {
		return err
	}
	mongoDatastore.client = *client
	err = mongoDatastore.client.Connect(context.Background())
	if err != nil {
		return err
	}

	//database := mongoDatastore.client.Database(config.DBName())
	database := mongoDatastore.client.Database("lolhub")
	mongoDatastore.db = *database

	// Access a MongoDB collection through a database
	events := database.Collection("events")
	addIndex(events, "event.id", -1, options.Index().SetUnique(true))

	frames := database.Collection("frames")
	addIndex(frames, "event_id", -1, nil)
	addIndex(frames, "game_id", -1, nil)

	// TODO: Remove
	databases, err := client.ListDatabaseNames(context.Background(), bson.M{})
	if err != nil {
		return err
	}
	fmt.Println(databases)

	return nil
}

func addIndex(col *mongo.Collection, field string, indexType interface{}, opt *options.IndexOptions) {
	mod := mongo.IndexModel{
		Keys: bson.M{
			field: indexType, // index in descending order
		},
		// create UniqueIndex option
		Options: opt,
	}
	ind, err := col.Indexes().CreateOne(context.Background(), mod)
	if err != nil {
		fmt.Println("Indexes().CreateOne() ERROR:", err)
		os.Exit(1) // exit in case of error
	} else {
		// API call returns string of the index name
		fmt.Println("CreateOne() index:", ind)
	}
}

// Now Returns a timestamp
func (mongoDatastore *mongoDatastore) Disconnect() {
	mongoDatastore.client.Disconnect(context.Background())
}

func (mongoDatastore *mongoDatastore) DB() *mongo.Database {
	return &mongoDatastore.db
}

func (mongoDatastore *mongoDatastore) Client() *mongo.Client {
	return &mongoDatastore.client
}

func (mongoDatastore *mongoDatastore) CreateIndex(collectionName string, field string, unique bool) bool {
	// 1. Lets define the keys for the index we want to create
	mod := mongo.IndexModel{
		Keys:    bson.M{field: 1}, // index in ascending order or -1 for descending order
		Options: options.Index().SetUnique(unique),
	}

	// 2. Create the context for this operation
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	// 3. Connect to the database and access the collection
	collection := mongoDatastore.db.Collection(collectionName)

	// 4. Create a single index
	_, err := collection.Indexes().CreateOne(ctx, mod)
	if err != nil {
		// 5. Something went wrong, we log it and return false
		fmt.Println(err.Error())
		return false
	}

	// 6. All went well, we return true
	return true
}
