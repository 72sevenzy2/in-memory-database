package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/72sevenzy2/in-memory-database/db"
)

// cli usage

// reusable func to display input format (SET and GET)

func displaySet() {
	fmt.Println("usage: SET <KeyName> <value>")
}

func displayGet() {
	fmt.Println("usage: GET <KeyName>")
}

func main() {
	b := db.NewDB()

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("> ")
		scanner.Scan()

		input := scanner.Text()
		parts := strings.Fields(input)

		if len(parts) == 0 {
			continue
		}

		UCinput := strings.ToUpper(parts[0])

		switch UCinput {
		case "SET":
			if len(parts) < 3 {
				fmt.Println("invalid SET format:")
				displaySet()
				continue
			}

			b.Set(parts[1], parts[2])
			fmt.Println("successful")
		case "GET":
			if len(parts) < 2 {
				fmt.Println("invalid GET format:")
				displayGet()
				continue
			}

			val, ok := b.Get(parts[1])
			if !ok {
				fmt.Println("key does not exist")
			} else {
				fmt.Println(val)
			}
		case "EXIT":
			fmt.Println("exited program")
			return

		default:
			fmt.Println("invalid command")
		}
	}

}
