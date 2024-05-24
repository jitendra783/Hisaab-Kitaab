package signup

import (
	signup "hisaab-kitaab/pkg/db/signUp"
	e "hisaab-kitaab/pkg/errors"
)

type User struct {
	User_ID string `json:"user_id"`
}
type ApiResponse struct {
	Data    []signup.User
	Status  bool      `json:"status"`
	Errors  []e.Error `json:"error"`
	Message string    `json:"message"`
}

type GetUserParam struct {
	Id string `uri:"id" binding:"required"`
}
