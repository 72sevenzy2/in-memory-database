package db

import (
	"encoding/binary"
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
	buf := make([]byte, 4)

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

// this will stay fixed as key will always be type string
func (v *DB) Del(key string) {
	delete(v.data, key) // built in delete() func to delete a particular key being held in the db map.
}
