package v1

import (
	"hisaab-kitaab/pkg/db"
	"hisaab-kitaab/pkg/service/v1/login"
	"hisaab-kitaab/pkg/service/v1/signup"
	"hisaab-kitaab/pkg/service/v1/user"

	"github.com/gin-gonic/gin"
)
//object
type srvObj struct {
	user.UserGroup
	login.LoginGroup
	signup.SignUpGroup
}

func NewServiceGroup(db db.DBLayer) ServiceLayer {
	return &srvObj{
		user.UserService(db),
		login.LoginService(db),
		signup.SignUpService(db),
	}
}
//interface
type ServiceLayer interface {
	user.UserGroup
	login.LoginGroup
	signup.SignUpGroup
	Status(*gin.Context)
	GetCurrentTime(c *gin.Context)
}
