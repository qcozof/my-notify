package model

//{"ok":false,"error_code":400,"description":"Bad Request: message text is empty"}

type TelegramRespModel struct {
	Ok          bool   `json:"ok"`
	ErrorCode   int    `json:"error_code"`
	Description string `json:"description"`
}
