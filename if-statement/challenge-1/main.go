package main

import (
	"fmt"
	"os"
)

func main() {
	username := "jack"
	password := "1888"

	args := os.Args
	if len(args) != 3 {
		fmt.Println(`Usage: [username] [password]`)
		return
	}

	if args[1] == username && args[2] == password {
		fmt.Println(`Access granted to "jack".`)
	} else if args[1] == "hacker" {
		fmt.Println(`Access denied for "hacker".`)
	} else {
		fmt.Println(`Invalid account or password.`)
	}

}
