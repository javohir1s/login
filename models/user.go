package models

type UserPrimaryKey struct {
	Id string `json:"id"`
	Login string `json:"login"`
}

type CreateUser struct {
	Name       string `json:"name"`
	Login      string `json:"login"`
	Password   string `json:"password"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type User struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	Login      string `json:"login"`
	Password   string `json:"password"`
	ExpiredAt string `json:"expires_at"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type UpdateUser struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	Login      string `json:"login"`
	Password   string `json:"password"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type GetListUserRequest struct {
	Offset int64  `json:"offset"`
	Limit  int64  `json:"limit"`
	Search string `json:"search"`
	Query  string `json:"query"`
}

type GetListUserResponse struct {
	Count int     `json:"count"`
	Users []*User `json:"users"`
}

type LoginRequest struct {
    Login    string `json:"login"`
    Password string `json:"password"`
}
