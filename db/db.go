package db

import (
	
)

// core logic
type DB struct {
	data map[string]string
}

func NewDB() *DB { // initialise a new map to hold data
	return &DB{
		data: make(map[string]string),
	}
}

func (v *DB) Set(key string, value string) error {
	v.data[key] = value

	return nil
}

func (v *DB) Get(key string) (string, bool) {
	val, ok := v.data[key]
	if !ok {
		return "", false
	}
	return val, true
}

func (v *DB) Del(key string) {
	delete(v.data, key) // built in delete() func to delete a particular key being held in the db map.
}
