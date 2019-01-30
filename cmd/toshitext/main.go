package main

import (
	"fmt"
	"time"
)

type Message struct {
	ID        int
	Message   string
	Date      time.Time
	Time      time.Time
	Sender    string
	Recipient string
	Spam      bool
}

func sendMessage(message string, sender string, recipient string, spam bool) {
	m := Message{Message: message, Sender: sender, Recipient: recipient, Spam: spam}
	fmt.Println(m.Message, m.Sender, m.Recipient, m.Spam)
}
func main() {
	fmt.Println("ToshiText Backend Api vAlpha")
}
