package db

import (
	"go.mongodb.org/mongo-driver/mongo"
	"lolesports/domain/db/champions"
	"lolesports/domain/db/events"
	"lolesports/domain/db/frames"
	"lolesports/infrastructure/dependencies"
	"lolesports/utils"
)

type Repository struct {
	EventsRepo events.Repository
	FramesRepo frames.Repository
	ChampsRepo champions.ChampionsRepository
}

type UseCase struct {
	Client     *mongo.Client
	Logger     utils.Handler
	Repository Repository
}

func New(container *dependencies.Container) *UseCase {

	repository := Repository{
		EventsRepo: events.New(container),
		FramesRepo: frames.New(container),
		ChampsRepo: champions.New(container),
	}

	useCase := &UseCase{
		Client:     container.MongoClient(),
		Logger:     container.Logger(),
		Repository: repository,
	}

	return useCase
}
