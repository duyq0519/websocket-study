package websocket

import (
	"bufio"
	"io"
	"net/http"
)

type Conn struct {
	config  *Config
	request *http.Request

	buf *bufio.ReadWriter
	rwc io.ReadWriteCloser
}

func newServerConn(rwc *io.ReadCloser, buf *bufio.ReadWriter, config *Config, handShake func(*Config, *http.Request) error) (conn *Conn, err error) {
	return nil, nil
}
