package main

import (
	"fmt"
	"time"

	sdk "github.com/official-stallion/go-sdk"
)

func main() {
	client, err := sdk.NewClient("st://localhost:9090")
	if err != nil {
		panic(err)
	}

	client.Subscribe("topic", func(data []byte) {
		fmt.Println(string(data))
	})

	client.Publish("topic", []byte("Hello"))

	time.Sleep(3 * time.Second)
}