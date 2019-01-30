package main

import "fmt"

type Message struct {
	ID          int
	Recipient   string
	MessageTime string
	Status      string
	Text        string
	FirstName   string
	LastName    string
	Country     string
	Sender      string
}

func sendMessage(message string, sender string, recipient string, spam bool) {
	m := Message{Message: message, Sender: sender, Recipient: recipient, Spam: spam}
	fmt.Println(m.Message, m.Sender, m.Recipient, m.Spam)
}

func main() {
	fmt.Println(m.Message, m.Sender, m.Recipient, m.Spam)
}
