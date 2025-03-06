package routes

import (
	"github.com/gin-gonic/gin"
	"go-automation-login/controllers"
)

// SetupRouter tạo các route cho API
func SetupRouter() *gin.Engine {
	router := gin.Default()

	// Định nghĩa API
	router.POST("/login", controllers.LoginHandler)
	router.POST("/telegram_webhook", controllers.TelegramWebhookHandler) // 📌 Webhook nhận phản hồi từ Telegram

	return router
}
