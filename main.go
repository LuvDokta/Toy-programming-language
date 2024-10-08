package main

import (
	"fmt"
	"log"
	"os"
	"os/user"

	"ksm/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		log.Fatal("Error: ", err)
	}

	fmt.Printf("Hello %s! This the Ksm Toy Programming language!\n", user.Username)
	fmt.Println()
	repl.StartRepl(os.Stdin, os.Stdout)
}
