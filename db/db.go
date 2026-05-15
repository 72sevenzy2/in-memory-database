package db

import (
	"encoding/binary"
	"errors"
)

// core logic
type DB struct {
	data map[string][]byte
}

func NewDB() *DB { // initialise a new map to hold data
	return &DB{
		data: make(map[string][]byte),
	}
}

func (v *DB) SetInt(key string, value uint32) error {
	buf := make([]byte, 4) // uint32's has a fixed memory size of 4 bytes

	binary.LittleEndian.PutUint32(buf, value)

	v.data[key] = buf
	return nil
}

func (v *DB) GetInt(key string) (uint32, bool) {
	data, ok := v.data[key]
	if !ok {
		return 0, false
	}

	value := binary.LittleEndian.Uint32(data)
	return value, true
}

// func to retrieve all values at once
func (v *DB) GetAllInt() (map[string]uint32, bool) {
	result := make(map[string]uint32)

	for k, v := range v.data {
		result[k] = binary.LittleEndian.Uint32(v) // serialize into uint32 type before assigning
	}
	if len(result) == 0 {
		return nil, false
	}
	return result, true
}

// string serialization and handlers

func (v *DB) SetString(key string, value string) (error, bool) {
	if value != "" {
		byteStr := []byte(value) // serialize string to type []byte, because db holds values of type []bte
		v.data[key] = byteStr
		return nil, true
	}
	return errors.New("please include a value aswell."), false
}

// get method for string

func (v *DB) GetString(key string) (string, bool) {
	val, ok := v.data[key];
	return string(val), ok
}

// display all string value data from db

func (v *DB) GetAllString() (map[string]string, bool) {
	results := make(map[string]string)

	for k, val := range v.data {
		results[k] = string(val)
	}

	if len(results) == 0 { // handling no existing data case
		return nil, false
	}

	return results, true
}

// this will stay fixed as key will always be type string
func (v *DB) Del(key string) {
	delete(v.data, key) // built in delete() func to delete a particular key being held in the db map.
}
