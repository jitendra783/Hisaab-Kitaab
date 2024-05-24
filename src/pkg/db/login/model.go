package login

import "time"

type Login struct {
	Id         string    `json:"userId" gorm:"primaryKey"`
	First_Name string    `json:"firstName"`
	Last_name  string    `json:"lastName"`
	Mobile_no  int64     `json:"mobileNumber"`
	Email_id   string    `json:"emailId"`
	Create_at  time.Time `json:"created_at"`
}
type LoginForm struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type FindForm struct {
	Email string `json:"email"`
}
type ResetForm struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
