package main

import (
	"encoding/json"
	"errors"
	"fmt"
)

type Message struct {
	Sender string `json:"sender"`
	Text   string `json:"text"`
}

var errEmptyMessage = errors.New("empty message")

func DecodeJSON(rawMsg string) (Message, error) {
	if len(rawMsg) == 0 {
		return Message{}, errEmptyMessage
	}

	msg := Message{}

	err := json.Unmarshal([]byte(rawMsg), &msg)
	if err != nil {
		return Message{}, fmt.Errorf("unmarshal: %w", err)
	}

	return msg, nil
}

func main() {
	msg, err := DecodeJSON("")
	if errors.Is(err, errEmptyMessage) {
		fmt.Println(msg, err)
	}

	msg, err = DecodeJSON("hello")

	if err != nil {
		fmt.Println(msg, err)
	}

	msg, err = DecodeJSON(`{"sender":"samir", "text":"Go,Go,Go"}`)
	fmt.Println(msg, err)
}
