package handlers

import (
	"log"
	"net/http"
)

type Goodbye struct {
	l *log.Logger
}

func NewGoodbye(l *log.Logger) *Goodbye {
	return &Goodbye{l}
}

// serverHTTP method

func (h *Goodbye) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Goodbye world from Service\n"))
}
