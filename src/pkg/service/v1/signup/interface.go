package signup

import (
	"hisaab-kitaab/pkg/db"

	"github.com/gin-gonic/gin"
)

type signUpObj struct{
	db db.DBLayer
}
type SignUpGroup interface {
	NewRegister(c *gin.Context)
}

func SignUpService(db db.DBLayer) SignUpGroup {
	return &signUpObj{db :db}
}
