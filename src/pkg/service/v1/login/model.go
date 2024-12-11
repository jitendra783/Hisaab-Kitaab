package login

import (
	dblogin "hisaab-kitaab/pkg/db/login"
	e "hisaab-kitaab/pkg/errors"
)

type User struct {
	User_ID string `json:"user_id"`
}
var Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type ApiResponse struct {
	Data    []dblogin.Login
	Status  bool      `json:"status"`
	Errors  []e.Error `json:"error"`
	Message string    `json:"message"`
}
type GetUserParam struct {
	Id string `uri:"id" binding:"required"`
}


