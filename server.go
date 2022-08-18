package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"lolesports/domain/api"
	"lolesports/domain/api/get"
	"lolesports/domain/db"
	"lolesports/infrastructure/application"
	"lolesports/infrastructure/configuration"
	"lolesports/infrastructure/dependencies"
	"os"
)

func main() {
	// scope
	scope := os.Getenv("SCOPE")
	if scope == "" {
		panic(fmt.Errorf("application initialization error - No scope defined"))
	}

	// build application context
	application.InitContext()

	// services initialization
	container := dependencies.StartDependencies()

	// configuration
	configuration := configuration.GetConfiguration()

	// db initialization
	mongoDS := container.MongoDataStore()
	err := mongoDS.Connect(configuration)
	if err != nil {
		panic(fmt.Errorf(err.Error()))
	}
	defer mongoDS.Disconnect()

	service := db.New(container)

	r := gin.Default()
	eventSyncer := api.NewEventSyncer(service)
	go eventSyncer.SyncLiveEvents(service)
	go eventSyncer.GetLiveEventsData(service)

	r.Use(cors.Default())

	group := r.Group("/lolesports")

	group.GET("/games", func(c *gin.Context) {
		get.LiveEvents(c, service)
	})

	group.GET("/event/:id", func(c *gin.Context) {
		get.Event(c, service)
	})

	group.GET("/frames/:gameId/:time", func(c *gin.Context) {
		get.Frames(c, service)
	}) // listen and serve on 0.0.0.0:8080

	r.Run()
}
