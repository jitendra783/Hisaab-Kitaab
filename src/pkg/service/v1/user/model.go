package user

import (
	e "hisaab-kitaab/pkg/errors"
)

type User struct {
	User_ID string `json:"user_id"`
}
type ApiResponse struct {
	Data    string     `json:"token"`
	Status  bool      `json:"status"`
	Errors  []e.Error `json:"error"`
	Message string    `json:"message"`
}

type GetUserParam struct{
	Id string `uri:"id" binding:"required"`
}
