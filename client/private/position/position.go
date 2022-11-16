package position

import (
	"github.com/rluisr/dydxv3/client/request"
)

type P struct {
	req *request.Request
}

func New(con Config) *P {
	{
		con.Verify()
	}

	return &P{
		req: con.Req,
	}
}
