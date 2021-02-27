package storage

import (
	"errors"
	"fmt"

	"github.com/philippgille/gokv"
	"github.com/philippgille/gokv/file"
)

type Storage struct {
	store gokv.Store
	dir   string
	key   string
}

func CreateStore(dir, key string) *Storage {

	options := file.DefaultOptions
	options.Directory = dir

	// Create client
	client, err := file.NewStore(options)
	if err != nil {
		panic(err)
	}
	defer client.Close()
	instance := &Storage{
		store: client,
		dir:   dir,
		key:   key,
	}
	return instance

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
