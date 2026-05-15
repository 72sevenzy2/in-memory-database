package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/72sevenzy2/in-memory-database/db"
	"github.com/72sevenzy2/in-memory-database/internal"
)

// cli usage

func main() {
	b := db.NewDB()

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("> ")
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
				internal.DisplaySet()
				continue
			}

			// parse the string given to type uint first
			f, err := strconv.ParseUint(parts[2], 10, 32)
			if err != nil {
				fmt.Println("invalid number")
				continue
			}

			b.SetInt(parts[1], uint32(f)) // then convert uint type to uint32
			fmt.Println("successful")
		case "GET":
			if len(parts) < 2 {
				fmt.Println("invalid GET format:")
				internal.DisplayGet()
				continue
			}

			val, ok := b.GetInt(parts[1])
			if !ok {
				fmt.Println("key does not exist")
			} else {
				fmt.Println(val)
			}
		case "FETCH":
			vals, ok := b.GetAllInt()
			if !ok {
				fmt.Println("no data available.")
			}
			for k, v := range vals {
				fmt.Println(k, v)
			}
		case "DEL":
			if len(parts) < 2 {
				internal.DisplayDel()
				continue
			}
			if _, ok := b.GetInt(parts[1]); ok {
				b.Del(parts[1])
				fmt.Println("successfully deleted key")
			} else {
				fmt.Println("key does not exist:", parts[1])
			}
		case "HELP":
			fmt.Println("General usage:")
			internal.DisplayGet()
			internal.DisplaySet()
			internal.DisplayDel()

			fmt.Println("\nTo exit: run <EXIT>")
		case "EXIT":
			fmt.Println("exited program")
			return

		default:
			fmt.Println("invalid command")
		}
	}

}
