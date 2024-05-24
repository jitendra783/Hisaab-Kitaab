package db

import (
	"database/sql"
	"hisaab-kitaab/pkg/db/login"
	signup "hisaab-kitaab/pkg/db/signUp"
	"hisaab-kitaab/pkg/db/user"

	"gorm.io/gorm"
)

type dbObj struct {
	user.UserGroup
	login.LoginGroup
	signup.SignUpGroup
}

func NewDbObj(psqlConn *gorm.DB, mysqlConn *sql.DB) DBLayer {
	temp := &dbObj{user.NewUserDBGroup(psqlConn),
		login.LoginDBGroup(psqlConn),
		signup.SignUpDBGroup(psqlConn),
	}
	return temp
}

type DBLayer interface {
	user.UserGroup
	login.LoginGroup
	signup.SignUpGroup
}
