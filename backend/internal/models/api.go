package models

type APIResponse[T any] struct {
	Success bool `json:"success"`
	Data    T    `json:"data,omitempty"`
	Error   any  `json:"error,omitempty"`
}

type LoginRequest struct {
	UserName string `form:"username" json:"username" xml:"username"  binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
}
type LoginRespBody struct {
	Token string `json:"token"`
}

func NewLoginRequest() *LoginRequest {

	return &LoginRequest{}

}
