package sender

import (
	"testing"
	"context"
	"github.com/aaronland/gomail/v2"		
)

func TestNullSender(t *testing.T) {

	ctx := context.Background()
	s, err := NewSender(ctx, "null://")

	if err != nil {
		t.Fatalf("Failed to create new null sender, %v", err)
	}
	
	msg := gomail.NewMessage()
	msg.SetBody("text/plain", "Hello world.")
	msg.SetHeader("From", "from@example.com")
	msg.SetHeader("To", "to@example.com")
	msg.SetHeader("Subject", "Null sender")
	
	err = gomail.Send(s, msg)

	if err != nil {
		t.Fatalf("Failed to send message, %v", err)
	}
	
}
