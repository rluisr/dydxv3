package trade

import (
	"github.com/rluisr/dydxv3/client/request"
)

type T struct {
	req *request.Request
}

func New(con Config) *T {
	{
		con.Verify()
	}

	return &T{
		req: con.Req,
	}
}
