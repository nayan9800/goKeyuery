package main

import (
	"fmt"

	"github.com/nayan9800/goKeyuery/server"
)

func main() {
	fmt.Println("goKeyuery")
	server.StartServer(":5000")

}
