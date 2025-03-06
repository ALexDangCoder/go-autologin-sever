package models

// LoginRequest là cấu trúc dữ liệu cho API login
type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
