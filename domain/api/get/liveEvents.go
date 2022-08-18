package get

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"lolesports/domain/db"
	"net/http"
)

func LiveEvents(c *gin.Context, service *db.UseCase) {
	limit := int64(10)
	findOpt := &options.FindOptions{
		Sort:  bson.D{{"event.startTime", -1}},
		Limit: &limit,
	}

	events, err := service.Repository.EventsRepo.Find(bson.M{}, findOpt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, events)
}
