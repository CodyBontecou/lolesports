package champions

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"lolesports/infrastructure/dependencies"
	"lolesports/utils"
)

// ChampionsRepository repository for Champions
type ChampionsRepository interface {
	Add(champion *Champion) error
	FindOne(filter bson.D) (*Champion, error)
	Find(filter bson.M, findOption *options.FindOptions) ([]Champion, error)
}

type championsRepository struct {
	logger utils.Handler
	db     *mongo.Database
}

// New constructor
func New(container *dependencies.Container) ChampionsRepository {
	mongoDB := container.MongoDataBase()
	logger := container.Logger()

	return &championsRepository{
		db:     mongoDB,
		logger: logger,
	}
}
