package login

import (
	"hisaab-kitaab/pkg/db"

	"github.com/gin-gonic/gin"
)

type loginObj struct {
	db db.DBLayer
}
type LoginGroup interface {
	Login(c *gin.Context)
	// ResetPassword(c *gin.Context)
}

func LoginService(db db.DBLayer) LoginGroup {
	return &loginObj{db: db}
}
