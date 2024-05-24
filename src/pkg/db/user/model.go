package user

import "time"

type User struct {
	Id       string    `json:"userId" gorm:"primaryKey"`
	Name     string    `json:"name"`
	Password string    `json:"password"`
	MobileNo int64     `json:"mobileNumber"`
	EmailId  string    `json:"emailId"`
	CreateAt time.Time `json:"created_at"`
}
type UserForm struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Mobile   int64  `json:"mobile"`
	Password string `json:"password"`
}
