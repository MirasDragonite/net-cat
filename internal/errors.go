package internal

import (
	"fmt"
	"net"
)

func errorHandler(conn net.Conn, s string, err error) {
	if err != nil {
		fmt.Fprint(conn, s)
	}
}
