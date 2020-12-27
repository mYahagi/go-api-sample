package main

import (
	"fmt"

	"github.com/mYahagi/go-api-sample/app/infrastructure"
)

func main() {
	fmt.Println("hello")
	db, err := infrastructure.DataBase{}.Connect()
	if err != nil {
		fmt.Printf("can not connect to database: %s", err)
		panic(err)
	}

	db.Close()
	fmt.Println("bye")
}
