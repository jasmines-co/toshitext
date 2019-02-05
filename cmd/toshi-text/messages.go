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
	Spam        bool
}

func sendMessage(text string, sender string, recipient string, spam bool) {
	m := Message{Text: text, Sender: sender, Recipient: recipient, Spam: spam}
	fmt.Println(m.Text, m.Sender, m.Recipient, m.Spam)
}

func main() {
	fmt.Println("sup yo!", "me", "you", true)
}
