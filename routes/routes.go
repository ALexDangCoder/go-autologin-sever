package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go-automation-login/controllers"
)

// SetupRouter tạo các route cho API
func SetupRouter() *gin.Engine {
	router := gin.Default()

	// Cấu hình CORS (Cho phép frontend gọi API)
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // Cho phép tất cả domain
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// Định nghĩa API
	router.POST("/login", controllers.LoginHandler)
	router.POST("/telegram_webhook", controllers.TelegramWebhookHandler)

	return router
}
