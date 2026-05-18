package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/72sevenzy2/in-memory-database/db"
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
				DisplaySet()
				continue
			}

			f, err := strconv.ParseUint(parts[2], 10, 32) // returns uint64, err.

			if err == nil { // its a int.
				err := b.SetInt(parts[1], uint32(f))
				if err != nil {
					fmt.Println(err.Error())
					continue
				} else {
					fmt.Println("successful.")
					continue
				}
			} else {
				// its a string if unable to parse to uint.
				err := b.SetString(parts[1], parts[2])
				if err != nil {
					fmt.Println(err.Error())
					continue
				} else {
					fmt.Println("successful.")
					continue
				}
			}
		case "GET":
			if len(parts) < 2 {
				fmt.Println("invalid GET format:")
				DisplayGet()
				continue
			}

			val, ok := b.GetInt(parts[1])
			if !ok {
				val2, ok2 := b.GetString(parts[1])
				if !ok2 {
					fmt.Println("data does not exist")
					continue
				} else {
					fmt.Println(val2)
					continue
				}
			} else {
				fmt.Println(val)
			}

		case "FETCH":
			vals, ok := b.GetAllInt() // returns map[string]uint32, bool
			vals2, ok2 := b.GetAllString() // returns map[string]string, bool

			if ok {
				for k, v := range vals {
					fmt.Println(k, int(v))
				}
			} else {
				fmt.Println("no data of value type: 'int' exist.")
			}
			if ok2 {
				for k, v := range vals2 {
					fmt.Println(k, v)
				}
			} else {
				fmt.Println("no data of value type: 'string' exists.")
			}

			// vals, ok := b.GetAllInt() // returns: map[string]uint32, bool
			// if !ok {
			// 	vals, ok := b.GetAllString()
			// 	if !ok {
			// 		fmt.Println("no existing data.")
			// 		continue
			// 	} else {
			// 		for k, v := range vals {
			// 			fmt.Println(k, v)
			// 		}
			// 	}
			// } else {
			// 	for k, v := range vals {
			// 		fmt.Println(k, v)
			// 	}
			// }
		case "DEL":
			if len(parts) < 2 {
				DisplayDel()
			}
			if _, ok := b.GetInt(parts[1]); ok {
				b.Del(parts[1])
				fmt.Println("successfully deleted key")
			} else {
				if _, ok := b.GetString(parts[1]); ok {
					b.Del(parts[1])
					fmt.Println("successfuly deleted key")
					continue
				} else {
					fmt.Println("key does not exit.", parts[1])
					continue
				}
			}
		case "HELP":
			fmt.Println("General usage:")
			DisplayGet()
			DisplaySet()
			DisplayDel()

			fmt.Println("\nTo exit: run <EXIT>")
		case "EXIT":
			fmt.Println("exited program")
			return

		default:
			fmt.Println("invalid command")
		}
	}

}
