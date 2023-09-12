package main

import (
	"fmt"
	"log"
	"os"

	"master/internal"
)

func main() {
	var port string
	arg := os.Args
	arg = arg[1:]
	if len(arg) == 0 {
		port = internal.StandardPort
	} else if len(arg) == 1 && internal.IsValidPort(arg[0]) {
		port = arg[0]
	} else {
		fmt.Println(internal.Usage)
		return
	}
	err := internal.NewServer(port)
	if err != nil {
		log.Fatal(err)
	}
}
