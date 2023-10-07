package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	server := args[0]
	dbName := args[1]
	fmt.Println(server)
	fmt.Println(dbName)
}
