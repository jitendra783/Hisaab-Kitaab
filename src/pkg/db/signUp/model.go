package signup

import "time"

type User struct {
	Id       string    `json:"Id"`
	Name     string    `json:"name"`
	Password string    `json:"password"`
	Mobile   int64     `json:"mobile"`
	Email    string    `json:"email"`
	CreateAt time.Time `json:"created_at"`
}
type SignupForm struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Mobile   int64  `json:"mobile"`
	Password string `json:"password"`
}
