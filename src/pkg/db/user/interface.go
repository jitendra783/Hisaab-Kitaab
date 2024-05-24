package user

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type userObj struct {
	psql *gorm.DB
}
type UserGroup interface {
	// UserDeleteByID(c *gin.Context) error
	// GetUserByID(c *gin.Context, id string) (User, error)
	// UpdateUserByID(c *gin.Context) (User, error)
	CreateUser(c *gin.Context, userinfo UserForm) (User, error)
}

func NewUserDBGroup(db *gorm.DB) UserGroup {
	return &userObj{psql: db}
}
