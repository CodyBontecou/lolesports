package frames

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"lolesports/infrastructure/dependencies"
	"lolesports/utils"
)

type Repository interface {
	Add(context context.Context, frame AllFrameData) (*AllFrameData, error)
	UpdateMany(filter bson.D, set bson.D) error
	Find(filter bson.M, findOptions *options.FindOptions) ([]*AllFrameData, error)
	FindOne(filter bson.M) (*AllFrameData, error)
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
