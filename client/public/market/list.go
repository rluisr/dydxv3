package market

import (
	"encoding/json"

	"github.com/rluisr/dydxv3/client/request"
	"github.com/xh3b4sd/tracer"
)

func (m *M) List(req ListRequest) (ListResponse, error) {
	var err error

	var byt []byte
	{
		byt, err = m.req.Get("markets", request.Values(req))
		if err != nil {
			return ListResponse{}, tracer.Mask(err)
		}
	}

	var res ListResponse
	{
		err = json.Unmarshal(byt, &res)
		if err != nil {
			return ListResponse{}, tracer.Mask(err)
		}
	}

	return res, nil
}
