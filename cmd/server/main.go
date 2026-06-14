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

				switch UCinput {
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
				}
			}
		}(conn)

	}

}
