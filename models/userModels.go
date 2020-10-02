package models

var RegisteredUsers []UserDetials

type UserDetials struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Mobile   string `json:"mobile"`
	Address  string `json:"address"`
}

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UpdatePassword struct {
	Username    string `json:"username"`
	OldPassword string `json:"oldPassword"`
	NewPassword string `json:"NewPassword"`
}
