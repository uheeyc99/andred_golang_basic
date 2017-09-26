package main

import (
	"broadcast"
	"time"
)

func main() {
	go broadcast.Response_Andrew()

	broadcast.Ask()
	time.Sleep(1000000000000)
}