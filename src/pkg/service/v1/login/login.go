package login

import (
	dblogin "hisaab-kitaab/pkg/db/login"
	e "hisaab-kitaab/pkg/errors"
	"hisaab-kitaab/pkg/logger"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (l *loginObj) Login(c *gin.Context) {
	logger.Log().Debug("start")
	defer logger.Log().Debug("end")
	var response ApiResponse
	var userinfo dblogin.LoginForm
	if err := c.BindJSON(&userinfo); err != nil {
		response.Errors = append(response.Errors, e.ErrorInfo[e.BadRequest].GetErrorDetails(""))
		c.JSON(http.StatusBadRequest, response)
		return
	}
	resp, err := l.db.Login(c, userinfo)
	if err != nil {
		logger.Log().Error("error ", zap.Error(err))
		response.Errors = append(response.Errors, e.ErrorInfo[err.Error()].GetErrorDetails(""))
		c.JSON(http.StatusBadRequest, response)
		return
	}
	logger.Log(c).Debug("user data", zap.Any("data", resp))
	response.Data = append(response.Data, resp)
	response.Status = true
	response.Message = "user logged in succefully"
	c.JSON(http.StatusOK, response)
}

// func (l *loginObj) ResetPassword(c *gin.Context) {
// 	logger.Log().Debug("start")
// 	defer logger.Log().Debug("end")
// 	var response ApiResponse
// 	var userinfo dblogin.LoginForm
// 	if err := c.BindJSON(&userinfo); err != nil {
// 		response.Errors = append(response.Errors, e.ErrorInfo[e.BadRequest].GetErrorDetails(""))
// 		c.JSON(http.StatusBadRequest, response)
// 		return
// 	}
// 	resp, err := l.db.ResetPassword(c, userinfo)
// 	if err != nil {
// 		logger.Log().Error("error ", zap.Error(err))
// 		response.Errors = append(response.Errors, e.ErrorInfo[e.AddDBError].GetErrorDetails(""))
// 		c.JSON(http.StatusBadRequest, response)
// 		return
// 	}

// 	response.Data = append(response.Data, resp)
// 	response.Status = true
// 	response.Message = "user Created succefully"
// 	c.JSON(http.StatusOK, response)
// }
