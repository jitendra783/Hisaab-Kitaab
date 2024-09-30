package signup

import (
	signup "hisaab-kitaab/pkg/db/signUp"
	e "hisaab-kitaab/pkg/errors"
	"hisaab-kitaab/pkg/logger"
	"hisaab-kitaab/pkg/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (r *signUpObj) NewRegister(c *gin.Context) {
	logger.Log().Debug("start")
	defer logger.Log().Debug("end")
	var response ApiResponse
	var userinfo signup.SignupForm
	if err := c.BindJSON(&userinfo); err != nil {
		response.Errors = append(response.Errors, e.ErrorInfo[e.BadRequest].GetErrorDetails(""))
		c.JSON(http.StatusBadRequest, response)
		return
	}
	// resp, err := r.db.Register(c, userinfo)
	// if err != nil {
	// 	logger.Log().Error("error ", zap.Error(err))
	// 	response.Errors = append(response.Errors, e.ErrorInfo[e.AddDBError].GetErrorDetails(""))
	// 	c.JSON(http.StatusBadRequest, response)
	// 	return
	// }
	token, _ := middleware.CreateToken(userinfo.Name)
	response.Data = token
	response.Status = true
	response.Message = "user Created succefully"
	c.JSON(http.StatusOK, response)
}
