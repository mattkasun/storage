package main

import (
	"fmt"

	"github.com/mattkasun/storage"
)

var users *storage.Storage

type User struct {
	Name string
	Pass string
}

func main() {
	users := storage.CreateStore("data", "users")

	userlist := []User{
		User{
			Name: "demo",
			Pass: "demo",
		},
		User{
			Name: "user",
			Pass: "fooo",
		},
	}

	if users.Save(userlist) {
		fmt.Println("userlist saved")
	} else {
		fmt.Println("failed to save userlist")
	}
	var user []User
	err := users.Get(&user)
	if err != nil {
		panic(err)
	}
	fmt.Printf("retrieved list %v %T\n", user, user)

	for i, item := range user {
		fmt.Println(i, item)
	}
}
