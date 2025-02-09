package user

import (
	"hisaab-kitaab/pkg/db"

	"github.com/gin-gonic/gin"
)

type userObj struct {
	db db.DBLayer
}

func UserService(db db.DBLayer) UserGroup {
	return &userObj{db: db}
}

type UserGroup interface {
	UserRegister(c *gin.Context)
	GetUserByID(c *gin.Context)
	// UserDetail(c *gin.Context)
	// UserDeleteByID(c *gin.Context)
	UserUpdate(c *gin.Context)
}
