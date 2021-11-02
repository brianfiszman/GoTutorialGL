package dtos

type GetUsersDTOResponse struct {
	Username string `json:"username"`
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
}

type GetUsersDTO struct {
	Username string `json:"username"`
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
}

type ResponseGetUsersDTO struct {
	Result []GetUsersDTO `json:"result"`
}
