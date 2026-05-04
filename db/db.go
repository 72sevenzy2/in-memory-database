package db

type DB struct {
	data map[string]string
}

func NewDB() *DB {
	return &DB{
		data: make(map[string]string),
	}
}

func (v *DB) Set(key string, value string) {
	v.data[key] = value
}

func (v *DB) Get(key string) (string, bool) {
	val, ok := v.data[key]
	if !ok {
		return "", false
	}
	return val, true
}

func (v *DB) Del(key string) {
	delete(v.data, key)
}