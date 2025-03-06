package main

import (
	"github.com/gin-gonic/gin" // Thêm package Gin
	"go-automation-login/routes"
	"go-automation-login/services"
	"log"
	"net/http"
)

// Middleware CORS để xử lý lỗi CORS
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		// Xử lý preflight request
		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}

func main() {
	// Chuyển Gin sang chế độ release
	gin.SetMode(gin.ReleaseMode) // 🔥 Tắt debug logs, tối ưu hiệu suất

	// Thiết lập Webhook Telegram khi server khởi động
	services.SetTelegramWebhook()

	// Khởi động Router
	r := routes.SetupRouter()

	// Thêm middleware CORS
	r.Use(CORSMiddleware())

	log.Println("🚀 Server đang chạy tại http://localhost:8080")
	r.Run(":8080")
}
