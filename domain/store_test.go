package domain

import (
	"testing"
)

func TestStore(t *testing.T) {
	var (
		key = "key"
		val = "value"
	)

	store := NewStore()

	err := store.Put(key, val)
	if err != nil {
		t.Errorf(NewExistKeyError(key).Error())
	}

	_, getError := store.Get(key)
	if getError != nil {
		t.Errorf(NewNotFoundError(key).Error())
	}
	err = store.Set(key, val)
	if err != nil {
		t.Errorf(NewNotFoundError(key).Error())
	}
}
