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

func Get(parts []string, b *db.DB, conn net.Conn) bool {
	if len(parts) < 2 || len(parts) > 2 {
		conn.Write([]byte("invalid GET format:\n"))
		conn.Write([]byte("GET <KeyName>\n"))
		return false
	}

	val, ok := b.GetInt(parts[1])
	if !ok {
		val2, ok2 := b.GetString(parts[1])
		if !ok2 {
			conn.Write([]byte("data does not exist\n"))
			return false
		}
		conn.Write([]byte(val2 + "\n"))
		return false
	}
	// fmt.Println(val)
	conn.Write([]byte(strconv.FormatUint(uint64(val), 10) + "\n")) // convert uint32 to readable format
	return true
}

func Fetch(b *db.DB, conn net.Conn) {
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
}
