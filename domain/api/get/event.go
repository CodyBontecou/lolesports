package get

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"lolesports/domain/db"
	"net/http"
)

func Event(c *gin.Context, service *db.UseCase) {
	pcID := c.Param("id")
	if pcID == "" {
		c.JSON(http.StatusBadRequest, "missing id")
		return
	}

	event, err := service.Repository.EventsRepo.FindOne(bson.M{"event.id": pcID})
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, event)
}
