package main

import (
	"fmt"
	"github.com/mYahagi/go-api-sample/app/infrastructure"
)

func main() {
	fmt.Println("hello")
	db := infrastructure.Connect()

	db.Close()
	fmt.Println("bye")
}
