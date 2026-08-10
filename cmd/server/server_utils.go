package main

import (
	"fmt"
	"math"
	"net"
	"strconv"

	"github.com/72sevenzy2/in-memory-database/db"
)

// utility functions for server/main.go

func Set(parts []string, conn net.Conn, b *db.DB) bool {
	if len(parts) < 3 || len(parts) > 3 {
		conn.Write([]byte("invalid SET format:\n"))
		conn.Write([]byte("SET <KeyName> <value>\n"))
		return false
	}

	f, err := strconv.ParseUint(parts[2], 10, 32) // returns uint64, err.

	if err == nil { // its a int.

		// prevent f from overflowing if number entered is too big
		if f > math.MaxUint32 {
			conn.Write([]byte("please include a number value less than unsigned int32.\n"))
			return false
		}

		err := b.SetInt(parts[1], uint32(f))
		if err != nil {
			fmt.Println(err.Error()) // print on server side
			return false
		}
		conn.Write([]byte("successful.\n"))
		return false
	}

	// its a string if unable to parse to uint.
	err2 := b.SetString(parts[1], parts[2])
	if err2 != nil {
		fmt.Println(err.Error())
		return false
	}
	return true
}
