package internal

import (
	"fmt"
	"net"
	"os"
)

func Art(conn net.Conn) {
	file, _ := os.ReadFile(pathToArt)
	fmt.Fprint(conn, welcomeMessage)
	for _, ch := range file {
		fmt.Fprint(conn, string(ch))
	}
	fmt.Fprint(conn, "\n")
}

func (c *Server) addUser(conn net.Conn, username string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.users[conn] = username
}

func (c *Server) DelUser(conn net.Conn) {
	c.mu.Lock()
	defer c.mu.Unlock()

	delete(c.users, conn)
}

func (c *Server) History(conn net.Conn, text string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.history += text + "\n"
}
