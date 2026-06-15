package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
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

				

		}
	}

}
