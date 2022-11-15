package main

import (
	"encoding/json"
	"fmt"
)

type permissions map[string]bool
type user struct {
	Name        string      `json:"username"`
	Password    string      `json:"-"`
	Permissions permissions `json:"perms,omitempty"`
}

func main() {
	users := []user{
		{"phong", "123456", permissions{"admin": true}},
		{"nga", "123456", permissions{"write": true}},
		{"nga", "123456", nil},
	}
	out, err := json.MarshalIndent(users, "", "\t")

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(out))
}
