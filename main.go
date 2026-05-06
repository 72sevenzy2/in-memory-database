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

		if parts[0] == "SET" || parts[0] == "set" {
			b.Set(parts[1], parts[2])
			fmt.Println("successful")
		}
		if parts[0] == "GET" || parts[0] == "get" {
			val, ok := b.Get(parts[1])
			if !ok {
				fmt.Println("key does not exist")
			} else {
				fmt.Println(val)
			}
		}
		if parts[0] == "EXIT" || parts[0] == "exit" {
			fmt.Println("exited program")
			return
		}
	}
}
