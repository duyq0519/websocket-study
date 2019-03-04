package websocket

type hybiServerHandshaker struct {
	*Config
	accept []byte
}
