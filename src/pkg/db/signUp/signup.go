package signup

import (
	e "hisaab-kitaab/pkg/errors"
	"hisaab-kitaab/pkg/logger"
	"log"

	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (l *signUpObj) Register(c *gin.Context, userinfo SignupForm) (User, error) {
	logger.Log().Debug("start")
	defer logger.Log().Debug("end")
	var user User
    var count int
	// var existingUser SignUp
	result := l.db.Where("email = ?", userinfo.Email).Table("mutualfundnew.user").Select("count(*)").Scan(&count)
	if result.Error != nil {
		log.Fatal(result.Error)
		if result.Error == gorm.ErrRecordNotFound {
			log.Println("creating a new record ")
			// Create a new user instance
			newUser := User{
				Email:    userinfo.Email,
				Password: userinfo.Password,
				Mobile:   userinfo.Mobile,
				Name:     userinfo.Name,
				Id:       "1",
				CreateAt: time.Now(),
			}

			// Insert the new user into the database
			createResult := l.db.Create(&newUser).Table("mutualfundnew.user")
			if createResult.Error != nil {
				log.Fatal(createResult.Error)
				return user, e.ErrorInfo[e.AddDBError]
			} else {
				user.Email = userinfo.Email
				user.Name = userinfo.Name
				user.Mobile = userinfo.Mobile
				return user, nil
			}
		} else {
			return user, e.ErrorInfo[e.GetDBError]
		}
	}
	user.Email = userinfo.Email
	user.Name = userinfo.Name
	user.Mobile = userinfo.Mobile
	return user, nil
}
