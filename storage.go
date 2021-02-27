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

func (s *Storage) Save(item interface{}) bool {
	//existing, err := s.Get()
	//fmt.Println(existing, err)
	//existing = append(existing, item)
	err := s.store.Set(s.key, item)
	if err != nil {
		return false
	}
	return true
}

func (s *Storage) Get() (interface{}, error) {
	value := new(interface{})
	found, err := s.store.Get(s.key, value)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	if !found {
		return nil, errors.New("not found")
	}
	return *value, nil
}
