package main

import (
	"errors"
	"log"
	"time"
)

var errorTCPCon = errors.New("ошибка соединения")

func SendTCP() error {
	return errorTCPCon
}

func DoHTTPCall() error {
	err := SendTCP()
	if err != nil {
		//существуют функции для проверки типов конкретных ошибок
		if errors.Is(err, errorTCPCon) {
			time.Sleep(1 * time.Second)
			return DoHTTPCall()
		}
		log.Println("unknown error on HTTP call", err)
	}

	return nil
}

func main() {
	var err = DoHTTPCall()
	if err != nil {
		log.Println("unknown error on HTTP call", err)
	}
}
