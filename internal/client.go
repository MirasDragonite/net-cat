package internal

import (
	"bufio"
	"errors"
	"fmt"
	"net"
	"time"
)

func connect(c *Server, conn net.Conn) {
	defer conn.Close()
	var userName string

	scan := bufio.NewScanner(conn)
	fmt.Fprint(conn, enterName)

	if scan.Scan() {
		userName = scan.Text()
		if isValidUserName(c, userName) {
			fmt.Fprint(conn, c.history)
			c.addUser(conn, userName)
		} else {
			errorHandler(conn, errorTextAboutInvalidText, errors.New(errorTextAboutInvalidText))
			connect(c, conn)
		}

	}

	defer c.DelUser(conn)
	status := Status{
		username: userName,
		conn:     conn,
		time:     time.Now(),
		text:     join,
	}
	c.History(conn, fmt.Sprintf(statusWriter, status.username, status.text))
	stsCh <- status
	defer func() {
		status := Status{
			username: userName,
			conn:     conn,
			time:     time.Now(),
			text:     left,
		}
		c.History(conn, fmt.Sprintf(statusWriter, status.username, status.text))
		stsCh <- status
	}()
	write(c, conn, userName, scan)
}

func write(c *Server, conn net.Conn, userName string, scan *bufio.Scanner) {
	for scan.Scan() {
		go func() {
			if isValidText(scan.Text()) {
				msg := Message{
					conn:     conn,
					username: userName,
					time:     time.Now(),
				}

				msg.text = scan.Text()
				c.History(conn, fmt.Sprintf(output, time.Now().Format(timeFormat), msg.username, msg.text))

				msgCh <- msg
			} else {
				fmt.Fprintf(conn, input, time.Now().Format(timeFormat), userName)
				write(c, conn, userName, scan)
			}
		}()
	}
}

func broadCast(c *Server) {
	for {
		select {
		case msg := <-msgCh:

			c.mu.Lock()
			for conn, name := range c.users {

				if msg.username != name {
					fmt.Fprint(conn, "\n", fmt.Sprintf(output, time.Now().Format(timeFormat), msg.username, msg.text), "\n")
				}

				fmt.Fprintf(conn, input, time.Now().Format(timeFormat), name)
			}
			c.mu.Unlock()

		case stat := <-stsCh:

			c.mu.Lock()
			for conn, name := range c.users {
				if stat.username != name {
					fmt.Fprint(conn, "\n", fmt.Sprintf(statusWriter, stat.username, stat.text), "\n")
				}
				fmt.Fprintf(conn, input, time.Now().Format(timeFormat), name)
			}
			c.mu.Unlock()
		}
	}
}
