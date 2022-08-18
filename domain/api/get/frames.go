package get

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"lolesports/domain/db"
	"net/http"
	"strconv"
)

func Frames(c *gin.Context, service *db.UseCase) {
	gameID := c.Param("gameId")
	if gameID == "" {
		c.JSON(http.StatusBadRequest, "missing gameID")
		return
	}

	timeRef := c.Param("time")
	if timeRef == "" {
		c.JSON(http.StatusBadRequest, "missing time")
		return
	}

	timeRefInt, err := strconv.Atoi(timeRef)
	if err != nil {
		c.JSON(http.StatusBadRequest, "time not integer")
		return
	}

	frames, err := service.Repository.FramesRepo.Find(bson.M{"game_id": gameID, "second": bson.M{"$gte": timeRefInt - 5, "$lte": timeRefInt + 10}}, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, frames)
}
