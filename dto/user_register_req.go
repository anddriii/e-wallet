package dto

type UserRegisterReq struct {
	FullName string `json"user_name"`
	Phone    string `json: "phone"`
	Email    string `json: "email"`
	UserName string `json: "username"`
	Password string `json: "password"`
}
