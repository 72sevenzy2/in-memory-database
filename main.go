package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/72sevenzy2/in-memory-database/db"
)

func main() {
	b := db.NewDB()

	scanner := bufio.NewScanner(os.Stdin)

	parts := strings.Split(scanner.Text(), " ")

	switch parts[0] {
	case "SET":
		b.Set(parts[1], parts[2])

	case "GET":
		val, ok := b.Get(parts[1])
		if ok {
			fmt.Println(val)
		} else {
			fmt.Println("key not found")
		}
	}
}
