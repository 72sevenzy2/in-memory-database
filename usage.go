package outer

import (
	"net"
)

// reusable functions to display input format (SET and GET)

func DisplaySet(conn net.Conn) {
	conn.Write([]byte("usage: SET <KeyName> <value>\n"));
}

func DisplayGet(conn net.Conn) {
	conn.Write([]byte("usage: GET <KeyName>\n"))
}

func DisplayDel(conn net.Conn) {
	conn.Write([]byte("usage: DEL <KeyName>\n"))
}