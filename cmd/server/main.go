package main

import (
	"fmt"
	"net"

	"github.com/72sevenzy2/in-memory-database/db"
)

func main() {
	// start tcp server
	ln, err := net.Listen("tcp", ":8080");
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

			buf := make([]byte, 1024) // preallocated buffer

			n, err := c.Read(buf)
			if err != nil {
				fmt.Println(err)
				return
			}

			fmt.Println("received:", string(buf[:n]))

		}(conn)
	
	}

}