package service

import (
	"hisaab-kitaab/pkg/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *serviceObj) Status(c *gin.Context) {
	logger.Log().Info("Status API called")
	c.String(http.StatusOK, "Working!")
}
