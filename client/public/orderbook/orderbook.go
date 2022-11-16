package orderbook

import (
	"github.com/rluisr/dydxv3/client/request"
)

type O struct {
	req *request.Request
}

func New(con Config) *O {
	{
		con.Verify()
	}

	return &O{
		req: con.Req,
	}
}
