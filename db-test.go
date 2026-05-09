package main

import (
	"fmt"
	"github.com/72sevenzy2/in-memory-database/db"
)

func TestDb() {
	b := *db.NewDB()
	b.SetInt("name", 100000)
	b.Del("name")

	val, ok := b.GetInt("name")
	if !ok {
		fmt.Println("not found", val)
	}
	fmt.Println(val)
}
