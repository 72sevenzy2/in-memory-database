package main

import (
	"fmt"
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
		b := db.NewDB() // db

		// seperate gorounine for each connection
		go func(c net.Conn) {
			defer c.Close() // close connection after reading

			buf := make([]byte, 1024) // preallocated buffer

			for {
				n, err := c.Read(buf)
				if err != nil {
					fmt.Println(err)
					return
				}

				input := string(buf[:n]) // n returns the number of bytes read

				parts := strings.Fields(input)

				if len(parts) == 0 { // avoid panic
					continue
				}

				// database logic

				UCinput := strings.ToUpper(parts[0]) // normalise to all capital

				switch UCinput {
				case "SET":
					ok := Set(parts, conn, b)
					if !ok {
						continue
					}
				case "GET":
					if len(parts) < 2 || len(parts) > 2 {
						conn.Write([]byte("invalid GET format:\n"))
						conn.Write([]byte("GET <KeyName>\n"))
						continue
					}

					val, ok := b.GetInt(parts[1])
					if !ok {
						val2, ok2 := b.GetString(parts[1])
						if !ok2 {
							conn.Write([]byte("data does not exist\n"))
							continue
						}
						conn.Write([]byte(val2 + "\n"))
						continue
					}
					// fmt.Println(val)
					conn.Write([]byte(strconv.FormatUint(uint64(val), 10) + "\n")) // convert uint32 to readable format

				case "FETCH":
					vals, ok := b.GetAllInt()      // returns map[string]uint32, bool
					vals2, ok2 := b.GetAllString() // returns map[string]string, bool

					if ok {
						for k, v := range vals {
							fmt.Fprintln(conn, k, int(v))
						}
					}
					if ok2 {
						for k, v := range vals2 {
							fmt.Fprintln(conn, k, v)
						}
					}

				case "DEL":
					if len(parts) < 2 || len(parts) > 2 {
						conn.Write([]byte("usage: DEL <KeyName>\n"))
					}
					if _, ok := b.GetInt(parts[1]); ok {
						b.Del(parts[1])
						conn.Write([]byte("successfully deleted key\n"))
					}
					if _, ok := b.GetString(parts[1]); ok {
						b.Del(parts[1])
						conn.Write([]byte("successfuly deleted key\n"))
						continue
					}
					fmt.Fprintln(conn, "key does not exit.", parts[1])
					continue

				case "HELP":
					conn.Write([]byte("General usage:\n"))
					usa.DisplayUsage(conn)

					conn.Write([]byte("\nto exit: run <exit\n"))
				case "EXIT":
					conn.Write([]byte("exited program.\n"))
					return

				default:
					conn.Write([]byte("invalid command.\n"))
				}

			}
		}(conn)

	}

}
