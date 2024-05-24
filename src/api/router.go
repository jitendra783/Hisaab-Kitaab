package api

import (
	"hisaab-kitaab/pkg/config"
	"hisaab-kitaab/pkg/service"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Router(obj service.ServiceGroupLayer, logger *zap.Logger) *gin.Engine {
	router := gin.New()
	router.Use(customLogger(logger))
	router.Use(gin.Recovery())
	router.GET("/health", obj.GetV1Service().Status)
	router.POST("/signup", obj.GetV1Service().NewRegister)
	router.POST("/login", obj.GetV1Service().Login)
	// router.POST("/reset", obj.GetV1Service().ResetPassword)

	// userGroup := router.Group("user")
	// {
	// 	// userGroup.POST("/register", obj.GetV1Service().UserRegister)
	// 	//  userGroup.GET("/getdetail/:id", obj.GetV1Service().GetUserByID)
	// 	// userGroup.PUT("/update", obj.GetV1Service().UserUpdate)
	// 	// userGroup.DELETE("/delete", obj.GetV1Service().UserDeleteByID)
	// }
	// signUpGroup := router.Group("user")
	// {
	// 	signUpGroup.POST("/register", obj.GetV1Service().UserRegister)
	// 	signUpGroup.GET("/getdetail/:id", obj.GetV1Service().GetUserByID)
	// 	signUpGroup.PUT("/update", obj.GetV1Service().UserUpdate)
	// 	signUpGroup.DELETE("/delete", obj.GetV1Service().UserDeleteByID)
	// }
	// loginGroup := router.Group("user")
	// {
	// 	loginGroup.POST("/register", obj.GetV1Service().UserRegister)
	// 	loginGroup.GET("/getdetail/:id", obj.GetV1Service().GetUserByID)
	// 	loginGroup.PUT("/update", obj.GetV1Service().UserUpdate)
	// 	loginGroup.DELETE("/delete", obj.GetV1Service().UserDeleteByID)
	// }
	return router
}

func customLogger(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		c.Next()

		if c.FullPath() != "/health" {
			latency := time.Since(start).Milliseconds()
			userID := c.GetString(config.USERID)
			uID := c.GetString(config.REQUESTID)
			// ucc := c.GetString(config.UCC)
			logger.Info(path,
				zap.String("requestID", uID),
				zap.String("leadId", "ucc"),
				zap.String("userId", userID),
				zap.Int("status", c.Writer.Status()),
				zap.String("method", c.Request.Method),
				zap.String("path", path),
				zap.String("query", query),
				zap.String("user-agent", c.Request.UserAgent()),
				zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
				zap.Int64("latency", latency),
			)
		}
	}
}
