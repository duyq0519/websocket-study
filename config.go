package websocket

import "net/url"

//Config ws 配置
type Config struct {
	Origin  url.URL
	Version int
}
