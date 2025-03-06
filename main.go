package main

import (
	"github.com/gin-gonic/gin" // Thêm package Gin
	"go-automation-login/routes"
	"go-automation-login/services"
	"log"
)

func main() {
	// Chuyển Gin sang chế độ release
	gin.SetMode(gin.ReleaseMode) // 🔥 Tắt debug logs, tối ưu hiệu suất

	// Thiết lập Webhook Telegram khi server khởi động
	services.SetTelegramWebhook()

	// Khởi động Router
	r := routes.SetupRouter()

	log.Println("🚀 Server đang chạy tại http://localhost:8080")
	r.Run(":8080")
}
