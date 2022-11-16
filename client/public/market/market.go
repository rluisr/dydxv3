package market

import (
	"github.com/rluisr/dydxv3/client/request"
)

type M struct {
	req *request.Request
}

func New(con Config) *M {
	{
		con.Verify()
	}

	return &M{
		req: con.Req,
	}
}
