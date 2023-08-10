package sender

import (
	"context"
	"testing"

	"github.com/aaronland/gomail/v2"
)

func TestStdoutSender(t *testing.T) {

	ctx := context.Background()
	s, err := NewSender(ctx, "stdout://")

	if err != nil {
		t.Fatalf("Failed to create new stdout sender, %v", err)
	}

	msg := gomail.NewMessage()
	msg.SetBody("text/plain", "Hello world.")
	msg.SetHeader("From", "from@example.com")
	msg.SetHeader("To", "to@example.com")
	msg.SetHeader("Subject", "Stdout sender")

	err = gomail.Send(s, msg)

	if err != nil {
		t.Fatalf("Failed to send message, %v", err)
	}

}
