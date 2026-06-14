package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"

	usa "github.com/72sevenzy2/in-memory-database"
	"github.com/72sevenzy2/in-memory-database/db"
)

// cli usage

func main() {
	conn, err := net.Dial("tcp", ":8080") // connect with tcp server
	if err != nil {
		panic(err)
	}

	b := db.NewDB()

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("> ")
		safe := scanner.Scan()
		if !safe {
			scanner.Err()
		}

		input := scanner.Text()
		parts := strings.Fields(input)

		if len(parts) == 0 {
			continue
		}

		UCinput := strings.ToUpper(parts[0])

		// send input to tcp server
		_, err := conn.Write([]byte(input))
		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		switch UCinput {

		case "GET":
			if len(parts) < 2 || len(parts) > 2 {
				fmt.Println("invalid GET format:")
				usa.DisplayGet()
				continue
			}

			val, ok := b.GetInt(parts[1])
			if !ok {
				val2, ok2 := b.GetString(parts[1])
				if !ok2 {
					fmt.Println("data does not exist")
					continue
				}
				fmt.Println(val2)
				continue
			}
			fmt.Println(val)

		case "FETCH":
			vals, ok := b.GetAllInt()      // returns map[string]uint32, bool
			vals2, ok2 := b.GetAllString() // returns map[string]string, bool

			if ok {
				for k, v := range vals {
					fmt.Println(k, int(v))
				}
			}
			if ok2 {
				for k, v := range vals2 {
					fmt.Println(k, v)
				}
			}
		case "DEL":
			if len(parts) < 2 || len(parts) > 2 {
				usa.DisplayDel()
			}
			if _, ok := b.GetInt(parts[1]); ok {
				b.Del(parts[1])
				fmt.Println("successfully deleted key")
			}
			if _, ok := b.GetString(parts[1]); ok {
				b.Del(parts[1])
				fmt.Println("successfuly deleted key")
				continue
			}
			fmt.Println("key does not exit.", parts[1])
			continue

		case "HELP":
			fmt.Println("General usage:")
			usa.DisplayGet()
			usa.DisplaySet()
			usa.DisplayDel()

			fmt.Println("\nto exit: run <exit")
		case "EXIT":
			fmt.Println("exited program")
			return

		default:
			fmt.Println("invalid command")
		}
	}

}
