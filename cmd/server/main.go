package main

import (
	"fmt"
	"math"
	"net"
	"strconv"
	"strings"

	usa "github.com/72sevenzy2/in-memory-database"
	"github.com/72sevenzy2/in-memory-database/db"
)

func main() {
	// start tcp server
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	for {
		conn, err := ln.Accept()

		if err != nil {
			fmt.Println(err.Error()) // err acceping connections
			continue
		}

		// seperate thread for each connection
		go func(c net.Conn) {
			defer c.Close() // close connection after reading

			b := db.NewDB() // db

			buf := make([]byte, 1024) // preallocated buffer

			for {
				n, err := c.Read(buf)
				if err != nil {
					fmt.Println(err)
					return
				}

				input := string(buf[:n])

				parts := strings.Fields(input)

				if len(parts) == 0 { // avoid panic
					continue
				}

				// database logic

				UCinput := strings.ToUpper(parts[0]) // normalise to all capital

				switch  UCinput{
				case "SET":
					if len(parts) < 3 || len(parts) > 3 {
						fmt.Println("invalid SET format:")
						usa.DisplaySet()
						continue
					}

					f, err := strconv.ParseUint(parts[2], 10, 32) // returns uint64, err.

					if err == nil { // its a int.

						// prevent f from overflowing if number entered is too big
						if f > math.MaxUint32 {
							fmt.Println("please include a number value less than unsigned int32.")
							continue
						}

						err := b.SetInt(parts[1], uint32(f))
						if err != nil {
							fmt.Println(err.Error())
							continue
						}
						fmt.Println("successful.")
						continue
					}
					// its a string if unable to parse to uint.
					err2 := b.SetString(parts[1], parts[2])
					if err2 != nil {
						fmt.Println(err.Error())
						continue
					}
					fmt.Println("successful.")
					continue

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
		}(conn)

	}

}
