package main

import (
	"first-app/src"
	"fmt"
)

func main() {
	u := &src.User{Name: "surya", Id: "1", Age: 24}
	user := src.NewUser(u)
	fmt.Println(user)
}
