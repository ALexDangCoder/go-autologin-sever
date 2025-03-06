package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

// Lấy Webhook URL từ biến môi trường (cấp bởi Render)
var webhookURL = os.Getenv("RENDER_WEBHOOK_URL") + "/telegram_webhook"

// SetTelegramWebhook - Đăng ký Webhook với Telegram
func SetTelegramWebhook() {
	url := fmt.Sprintf("%s/setWebhook", apiBaseURL)

	payload := map[string]string{
		"url": webhookURL,
	}
	jsonPayload, _ := json.Marshal(payload)

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		log.Printf("❌ Lỗi khi thiết lập Webhook: %v", err)
		return
	}
	defer resp.Body.Close()

	var result struct {
		Ok          bool   `json:"ok"`
		Description string `json:"description"`
	}

	json.NewDecoder(resp.Body).Decode(&result)

	if result.Ok {
		log.Println("✅ Webhook Telegram đã được thiết lập thành công!")
	} else {
		log.Printf("⚠️ Lỗi thiết lập Webhook: %s", result.Description)
	}
}
