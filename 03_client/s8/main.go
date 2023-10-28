package main

import (
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("main.go")
	if err != nil {
		panic(err)
	}
	resp, err := http.Post("http://localhost:18888", "text/plain", file)
	if err != nil {
		// 送信失敗
		panic(err)
	}
	log.Println("Status:", resp.Status)

	reader := strings.NewReader("テキスト")
	resp, err = http.Post("http://localhost:18888", "text/plain", reader)
	if err != nil {
		// 送信失敗
		panic(err)
	}
	log.Println("Status:", resp.Status)
}
