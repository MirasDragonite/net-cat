package internal

import (
	"net"
	"os"
	"sync"
	"time"
)

var (
	msgCh = make(chan Message)
	stsCh = make(chan Status)
	c     = make(chan os.Signal, 1)
)

type Message struct {
	username string
	text     string
	time     time.Time
	conn     net.Conn
}

type Status struct {
	username string
	time     time.Time
	conn     net.Conn
	text     string
}

type Server struct {
	mu      sync.Mutex
	users   map[net.Conn]string
	history string
}
