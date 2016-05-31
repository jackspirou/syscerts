package autoload

import (
	"crypto/tls"
	"net"
	"net/http"
	"time"

	"github.com/jackspirou/syscerts"
)

func init() {
	http.DefaultTransport = &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		Dial: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		}).Dial,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
		TLSClientConfig:       &tls.Config{RootCAs: syscerts.SystemRootsPool()},
	}
}
