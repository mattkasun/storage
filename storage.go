package storage

import (
	"errors"

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
	err := s.store.Set(s.key, item)
	if err != nil {
		return false
	}
	return true
}

func (s *Storage) Get(value interface{}) error {
	found, err := s.store.Get(s.key, value)
	if err != nil {
		return err
	}
	if !found {
		return errors.New("not found")
	}
	return nil
}
