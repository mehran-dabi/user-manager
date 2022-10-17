package controller

type CreateRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	NickName  string `json:"nick_name"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	Country   string `json:"country"`
}
