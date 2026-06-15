package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

// cli usage

func main() {
	conn, err := net.Dial("tcp", ":8080") // connect with tcp server
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("> ")
		safe := scanner.Scan()
		if !safe {
			scanner.Err()
		}

		input := scanner.Text()

		// send input to tcp server
		_, err := conn.Write([]byte(input))
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
	}

}
