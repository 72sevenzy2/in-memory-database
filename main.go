package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/72sevenzy2/in-memory-database/db"
)

// cli usage

func main() {
	b := db.NewDB()

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("> ")
		scanner.Scan()

		input := scanner.Text()
		parts := strings.Split(input, " ")

		switch parts[0] {
		case "SET":
			b.Set(parts[1], parts[2])
			fmt.Println("successful")

		case "GET":
			val, ok := b.Get(parts[1])
			if !ok {
				fmt.Println("key does not exist")
			}
			fmt.Println(val)

		case "EXIT":
			fmt.Println("exited program")
			return
		}
	}
}
