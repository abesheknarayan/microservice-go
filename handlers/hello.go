package handlers

import (
	"log"
	"net/http"
)

type Hello struct {
	l *log.Logger
}

func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

// serverHTTP method

func (h *Hello) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Hello world from Service\n"))
}
