package main

import (
	"fmt"
	"github.com/tantoni228/server_calculator/db/conection"
	"github.com/tantoni228/server_calculator/http/server"
)

func main() {
	err := conection.CreateDatabase()
	if err != nil {
		fmt.Println(err)
	}
	server.Run()

}
