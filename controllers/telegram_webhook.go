package controllers

import (
	"github.com/gin-gonic/gin"
	"go-automation-login/services"
	"log"
	"net/http"
	"strings"
)

// Cấu trúc dữ liệu từ Telegram
type TelegramUpdate struct {
	UpdateID      int `json:"update_id"`
	CallbackQuery *struct {
		ID   string `json:"id"`
		Data string `json:"data"`
	} `json:"callback_query"`
}

// TelegramWebhookHandler - Nhận phản hồi từ Telegram
func TelegramWebhookHandler(c *gin.Context) {
	var update TelegramUpdate

	// Kiểm tra dữ liệu đầu vào
	if err := c.ShouldBindJSON(&update); err != nil {
		log.Printf("❌ Lỗi nhận phản hồi từ Telegram: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dữ liệu không hợp lệ"})
		return
	}

	// Kiểm tra xem có CallbackQuery không
	if update.CallbackQuery == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Không có callback từ Telegram"})
		return
	}

	// Xử lý phản hồi từ Telegram
	callbackData := update.CallbackQuery.Data
	parts := strings.Split(callbackData, "|")
	if len(parts) != 4 || parts[0] != "status" {
		log.Println("⚠️ Phản hồi không đúng định dạng:", callbackData)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Phản hồi không đúng định dạng"})
		return
	}

	status := parts[2]
	requestID := parts[3]

	// Gửi phản hồi vào kênh đúng requestID
	services.HandleTelegramResponse(requestID, status)

	// Trả về phản hồi thành công cho Telegram
	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "Phản hồi đã được xử lý!"})
}
