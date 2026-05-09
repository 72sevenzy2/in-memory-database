package internal

import (
	"fmt"
)

// reusable functions to display input format (SET and GET)

func DisplaySet() {
	fmt.Println("usage: SET <KeyName> <value>")
}

func DisplayGet() {
	fmt.Println("usage: GET <KeyName>")
}

func DisplayDel() {
	fmt.Println("usage: DEL <KeyName>")
}