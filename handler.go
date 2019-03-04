package websocket

import (
	"fmt"
	"net/http"
)

type Handler func(*Conn)

func (h Handler) ServerHTTP(w http.ResponseWriter, r *http.Request) {
	s := Server{Handler: h, HandShake: CheckOrigin}
	s.serveWebSocket(w, r)
}

//func(*Config, *http.Request) 这个方法有什么意思？
func CheckOrigin(config *Config, r *http.Request) error {
	origin, err := Origin(config, r)
	if err == nil && origin == nil {
		return fmt.Errorf("null origin")
	}
	return err
}
