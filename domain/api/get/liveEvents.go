package get

import (
	"github.com/gin-gonic/gin"
	"lolesports/domain/api"
	"net/http"
)

func LiveEvents(es *api.EventSyncer, c *gin.Context) {
	es.Mu.Lock()
	c.JSON(http.StatusOK, es.LiveEvents)
	es.Mu.Unlock()
}
