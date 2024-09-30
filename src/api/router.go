package api

import (
	"hisaab-kitaab/pkg/middleware"
	"hisaab-kitaab/pkg/service"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Router(obj service.ServiceGroupLayer, logger *zap.Logger) *gin.Engine {
	router := gin.New()
	router.Use(middleware.CustomLogger(logger))
	router.Use(gin.Recovery())
	router.GET("/health", obj.GetV1Service().Status)
	router.POST("/signup", obj.GetV1Service().NewRegister)
	router.POST("/login", obj.GetV1Service().Login)
	// router.POST("/forgot", obj.GetV1Service().UserRegister)
	// router.POST("/reset", obj.GetV1Service().UserRegister)
	router.Use(middleware.AuthMiddleware())

	userGroup := router.Group("user")
	{
		userGroup.GET("/getdetail/:id", obj.GetV1Service().GetUserByID)
		userGroup.PUT("/update", obj.GetV1Service().UserRegister)
		// userGroup.DELETE("/deregister", obj.GetV1Service().UserRegister)
	}
	return router
}
