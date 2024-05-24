package signup

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type signUpObj struct {
	db *gorm.DB
}

func SignUpDBGroup(psqlConn *gorm.DB) SignUpGroup {
	return &signUpObj{db: psqlConn}
}

type SignUpGroup interface {
	Register(c *gin.Context, userinfo SignupForm) (User, error)
}
