package main

import (
	"fmt"

	"github.com/72sevenzy2/in-memory-database/db"
)

func main() {
	b := *db.NewDB()
	b.Set("name", "sam")
	b.Del("name")

	val, ok := b.Get("name")
	if !ok {
		fmt.Println("not found", val)
	}
	fmt.Println(val)
}
