package outer

import (
	"net"
)

// reusable functions to display input format (SET and GET)

func DisplayUsage(conn net.Conn) {
	conn.Write([]byte(`
		SET <KeyName> <value>
		GET <KeyName>
		DEL <KeyName>
	\n`))
}
