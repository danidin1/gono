package main

import (
	"fmt"

	log "github.com/danidin1/gono/logger"
)

func Add(a, b int) int {
	return a + b
}

func main() {

	answer := Add(4, 5)
	fmt.Println(answer)
	fmt.Println(log.GetLogName())
}
