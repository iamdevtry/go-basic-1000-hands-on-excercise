package main

import (
	"fmt"
	"os"
)

const (
	usage    = "Usage: [username] [password]"
	errUser  = "Access denied for %q.\n"
	errPwd   = "Invalid password for %q.\n"
	accessOK = "Access granted to %q.\n"

	user1, user2        = "jack", "inanc"
	password, password2 = "1888", "1879"
)

func main() {
	args := os.Args
	if len(args) != 3 {
		fmt.Println(usage)
		return
	}

	u, p := args[1], args[2]

	if u != user1 && u != user2 {
		fmt.Printf(errUser, u)
	} else if u == user1 && p == password {
		fmt.Printf(accessOK, u)
	} else if u == user2 && p == password2 {
		fmt.Printf(accessOK, u)
	} else {
		fmt.Printf(accessOK, u)
	}
}
