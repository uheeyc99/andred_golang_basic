package main

import (
	"fmt"
	"time"
	"detect"
)

func main() {
	fmt.Println(time.Now())
	detect.Detect()
}