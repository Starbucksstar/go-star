package entity

type AuthUser struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	Id       string `json:"id"`
}
