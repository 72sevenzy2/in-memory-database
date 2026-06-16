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
	reader := bufio.NewReader(conn)

	for {
		fmt.Print("> ")
		safe := scanner.Scan()
		if !safe {
			scanner.Err()
			break
		}

		input := scanner.Text()

		// send input to tcp server
		_, err := conn.Write([]byte(input))
		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		resp, err2 := reader.ReadString('\n')
		if err2 != nil {
			fmt.Println(err2.Error())
			continue
		}

		fmt.Println(resp)
	}

}
