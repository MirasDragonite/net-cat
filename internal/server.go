package internal

import (
	"errors"
	"fmt"
	"net"
	"os"
	"os/signal"
	"sync"
)

func NewServer(port string) error {
	ln, err := net.Listen(typeOfNetwork, ":"+port)
	fmt.Printf(messageAboutStart, port)
	if err != nil {
		return err
	}
	defer ln.Close()

	chat := &Server{
		mu:    sync.Mutex{},
		users: map[net.Conn]string{},
	}

	signal.Notify(c, os.Interrupt)

	go func() {
		for sig := range c {
			if sig == os.Interrupt {
				chat.mu.Lock()
				defer chat.mu.Unlock()
				os.WriteFile(pathToLogFile, []byte(chat.history), 0o777)
				os.Exit(0)
			}
		}
	}()

	go broadCast(chat)
	for {
		conn, err := ln.Accept()
		if err != nil {
			os.Exit(0)
		}
		chat.mu.Lock()

		if len(chat.users) < 10 {
			Art(conn)
			go connect(chat, conn)
		} else {
			errorHandler(conn, errorTextAboutTooMuchUsers, errors.New(errorTextAboutTooMuchUsers))
			conn.Close()
		}
		chat.mu.Unlock()
	}
}
