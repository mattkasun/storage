package user

import (
	"errors"
	"fmt"

	"github.com/philippgille/gokv"
	"github.com/philippgille/gokv/file"
)

type Users []User

type User struct {
	Name string
	Pass string
}

var store gokv.Store

func init() {
	fmt.Println("vim-go")

	options := file.DefaultOptions
	options.Directory = "data"

	// Create client
	client, err := file.NewStore(options)
	if err != nil {
		panic(err)
	}
	store = client
	defer client.Close()

}

func GetUsers() (Users, error) {
	var users Users
	found, err := store.Get("Users", &users)
	if err != nil {
		fmt.Println("error in GetUser")
		return Users{}, err
	}
	if !found {
		return Users{}, errors.New("Users not found")
	}
	return users, nil

}
