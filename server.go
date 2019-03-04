package websocket

import "net/http"

type Server struct {
	Config
	HandShake func(*Config, *http.Request) error
	Handler
}

func (s *Server) serveWebSocket(w http.ResponseWriter, r *http.Request) {
	rwc, buf, err := w.(http.Hijacker).Hijack()
	if err != nil {
		panic("Hijack failed:" + err.Error())
	}
	defer rwc.Close()
	

}
