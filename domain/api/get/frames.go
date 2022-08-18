package get

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
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

	if len(frames) == 0 {
		limit := int64(1)
		findOpt := &options.FindOptions{
			Sort:  bson.D{{"second", -1}},
			Limit: &limit,
		}

		framesAux, err := service.Repository.FramesRepo.Find(bson.M{"game_id": gameID}, findOpt)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		if len(framesAux) > 0 {
			if timeRefInt > framesAux[0].Second {
				framesAux[0].Second = timeRefInt
				frames = framesAux
			}
		}
	}

	c.JSON(http.StatusOK, frames)
}
