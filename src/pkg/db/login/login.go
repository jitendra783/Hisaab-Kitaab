package login

import (
	"errors"
	e "hisaab-kitaab/pkg/errors"
	"hisaab-kitaab/pkg/logger"
	"log"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func (l *loginObj) Login(c *gin.Context, userinfo LoginForm) (Login, error) {
	logger.Log().Debug("start")
	defer logger.Log().Debug("end")
	var user Login
	// query := "select count(*) from mutualfundnew.user where email =? and password =? limit 1"
	// logger.Log().Debug("query", zap.Any("q", query))
	// Execute the query
	if err := l.db.Find(&user, "user.email= ? and  user.password =?", userinfo.Username, userinfo.Password).Error; err != nil {
		if errors.Is(err, gorm.ErrInvalidDB) {
			log.Fatal("Failed to execute the query:", err)
			logger.Log().Debug("query", zap.Error(err))
			return user, e.ErrorInfo[e.GetDBError]
		} else if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Fatal("record not found", err)
			return user, e.ErrorInfo[e.NoDataFound]
		} else {
			log.Fatal("record not found", err)
			return user, e.ErrorInfo[e.BadRequest]
		}
	}
	return user, nil
}

// func (l *loginObj) ResetPassword(c *gin.Context, userinfo LoginForm) (Login, error) {
// 	logger.Log().Debug("start")
// 	defer logger.Log().Debug("end")
// 	//id := c.GetString(config.USERID)
// 	var user Login
// 	query := "update table hisaab_kitaab.user set password = ? where email = ? "
// 	// Execute the query
// 	_, err := l.db.Exec(query, userinfo.Password, userinfo.Username)
// 	logger.Log().Debug("query", zap.Any("q", query))
// 	if err != nil {
// 		log.Fatal("Failed to execute the query:", err)
// 		return user, err
// 	}
// 	//undefined error while checking in db
// 	// user, err = u.GetUserByID(c,id)
// 	// if err != nil{
// 	// 	log.Fatal("Failed to execute the query:", err)
// 	// 	return user, err
// 	// }
// 	return user, nil
// }
