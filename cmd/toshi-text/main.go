package main

import (
	"fmt"
	"time"
)

type TextMessage struct {
	ID        int
	Message   string
	Date      time.Time
	Time      time.Time
	Sender    string
	Recipient string
}

func sendMessage(message string, sender string, recipient string) {
	m := TextMessage{Message: message, Sender: sender, Recipient: recipient}
	fmt.Println(m.Message, m.Sender, m.Recipient)
}
func main() {
	fmt.Println("ToshiText Backend Api vAlpha")
}
