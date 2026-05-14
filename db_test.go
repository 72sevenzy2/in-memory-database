package main

import (
	"fmt"
	"github.com/72sevenzy2/in-memory-database/db"
	"testing"
)

func TestDb(t *testing.T) {
	b := *db.NewDB()
	b.SetInt("name", 100000)
	b.Del("name")

	val, ok := b.GetInt("name")
	if !ok {
		fmt.Println("not found", val)
	}
	fmt.Println(val)
}
