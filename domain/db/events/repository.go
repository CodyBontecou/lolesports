package events

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	api "lolesports/domain/service"
	"lolesports/infrastructure/dependencies"
	"lolesports/utils"
)

type Repository interface {
	Add(context context.Context, account AllEventData) (*AllEventData, error)
	UpdateMany(filter bson.D, set bson.D) error
	Find(filter bson.M, findOptions *options.FindOptions) ([]*AllEventData, error)
	FindOne(filter bson.M) (*AllEventData, error)
	AddFrames2Event(ctx context.Context, eventID string, lastFrame int, frames []api.Frame) error
	AddWindowFrames2Event(ctx context.Context, eventID string, lastFrame int, frames []api.WindowFrame) error
	GetNotFetched() ([]*AllEventData, error)
}

type scenesRepository struct {
	logger utils.Handler
	db     *mongo.Database
}

// New constructor
func New(container *dependencies.Container) Repository {
	mongoDB := container.MongoDataBase()
	logger := container.Logger()

	return &scenesRepository{
		db:     mongoDB,
		logger: logger,
	}
}
