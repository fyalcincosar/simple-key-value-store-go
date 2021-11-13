package domain

import (
	"fmt"
)

type Store struct {
	db map[string]string
}

func NewStore() (storeDB *Store) {
	store := &Store{db: map[string]string{}}
	return store
}
func (s Store) Get(key string) (string, error) {

	value := s.db[key]

	if len(value) > 0 {
		return value, nil
	} else {
		return "", NewNotFoundError(key)
	}
}
func (s Store) Set(key string, value string) error {
	_, err := s.Get(key)

	if err != nil {
		return err
	} else {
		s.db[key] = value
		return nil
	}
}
func (s Store) Put(key string, value string) error {
	_, err := s.Get(key)

	if err != nil {
		s.db[key] = value
		return nil
	} else {
		return NewExistKeyError(key)
	}
}
func (s Store) GetDb() map[string]string {
	return s.db
}

type NotFoundError struct {
	nonexistentKey string
}

func NewNotFoundError(nonexistentKey string) error {
	return &NotFoundError{nonexistentKey}
}

func (n *NotFoundError) Error() string {
	return fmt.Sprintf("Key: %s not found!", n.nonexistentKey)
}

type ExistKeyError struct {
	existKey string
}

func NewExistKeyError(existKey string) error {
	return &ExistKeyError{existKey}
}
func (e *ExistKeyError) Error() string {
	return fmt.Sprintf("Exist Key: %s!", e.existKey)
}
