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

func (g *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("hey"))
}
