package login

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type loginObj struct {
	db *gorm.DB
}

func LoginDBGroup(psql *gorm.DB) LoginGroup {
	return &loginObj{db: psql}
}

type LoginGroup interface {
	Login(c *gin.Context, userinfo LoginForm) (Login, error)
	// ResetPassword(c *gin.Context, userinfo LoginForm) (Login, error)
}
