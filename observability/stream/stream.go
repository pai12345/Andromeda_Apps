package stream

import "fmt"

func (msg CStream) SendMessage(message string, channel chan string) {
	result := fmt.Sprintf("Message %v sent from %v -> %v", message, msg.Source, msg.Destination)
	channel <- result
}

func (msg WStream) SendMessage(message string) {
	fmt.Printf("Message %v sent", message)
}
