package notify

import (
	"fmt"
	"testing"
)

func TestEmail_Send(t *testing.T) {
	InitConfig("")

	var email Email
	res, err := email.Send()
	if err == nil {
		fmt.Println(err)
	}
	fmt.Println(res)
}

func TestTelegram(t *testing.T) {
	InitConfig("")

	Telegram("content")
}
