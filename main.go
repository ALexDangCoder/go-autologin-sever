package main

import (
	"go-automation-login/routes"
	"go-automation-login/services"
	"log"
)

func main() {
	// Thiết lập Webhook Telegram khi server khởi động
	services.SetTelegramWebhook() // ✅ Gọi thiết lập webhook

	// Khởi động Router
	r := routes.SetupRouter()

	log.Println("🚀 Server đang chạy tại http://localhost:8080")
	r.Run(":8080")
}
