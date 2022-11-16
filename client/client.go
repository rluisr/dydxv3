package client

import (
	"net"
	"net/http"
	"time"

	"github.com/rluisr/dydxv3/client/private"
	"github.com/rluisr/dydxv3/client/public"
	"github.com/rluisr/dydxv3/client/request"
	"github.com/rluisr/dydxv3/client/secret"
)

type Client struct {
	Pri *private.Private
	Pub *public.Public
}

func New(con Config) *Client {
	con.Verify()

	httpTransport := &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   3 * time.Second,
			KeepAlive: 24 * time.Hour,
		}).DialContext,
		MaxIdleConns:          0,
		MaxIdleConnsPerHost:   100,
		IdleConnTimeout:       24 * time.Hour,
		TLSHandshakeTimeout:   3 * time.Second,
		ResponseHeaderTimeout: 10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}

	var pri *private.Private
	if con.Sec != (secret.Secret{}) {
		pri = private.New(private.Config{
			Req: request.New(request.Config{
				Cli: &http.Client{
					Transport: httpTransport,
					Timeout:   2 * time.Second,
				},
				Pri: true,
				Sec: con.Sec,
				Tes: con.Tes,
			}),
			Sec: con.Sec,
			Tes: con.Tes,
		})
	}

	var pub *public.Public
	pub = public.New(public.Config{
		Req: request.New(request.Config{
			Cli: &http.Client{
				Transport: httpTransport,
				Timeout:   5 * time.Second,
			},
			Tes: con.Tes,
		}),
		Tes: con.Tes,
	})

	return &Client{
		Pri: pri,
		Pub: pub,
	}
}
