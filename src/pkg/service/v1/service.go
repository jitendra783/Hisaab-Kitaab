package v1

import (
	"hisaab-kitaab/pkg/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *srvObj) Status(c *gin.Context) {
	c.String(http.StatusOK, "Working!")
}

func (s *srvObj) GetCurrentTime(c *gin.Context) {
	logger.Log(c).Debug("start")
	defer logger.Log(c).Debug("end")

}
