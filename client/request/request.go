package request

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/rluisr/dydxv3/client/secret"
	"github.com/rluisr/dydxv3/errors"
	"github.com/xh3b4sd/tracer"
)

type Request struct {
	cli *http.Client
	pri bool
	sec secret.Secret
	tes bool
}

func New(con Config) *Request {
	{
		con = con.Ensure()
	}

	{
		con.Verify()
	}

	return &Request{
		cli: con.Cli,
		pri: con.Pri,
		sec: con.Sec,
		tes: con.Tes,
	}
}

func (r *Request) endp(pat string) string {
	var end string
	if r.tes {
		end = "https://api.stage.dydx.exchange"
	} else {
		end = "https://api.dydx.exchange"
	}

	return fmt.Sprintf("%s/%s/%s", end, "v3", pat)
}

func (r *Request) request(met string, pat string, dat interface{}) ([]byte, error) {
	var (
		b   []byte
		err error
	)

	if dat != nil {
		b, err = json.Marshal(dat)
		if err != nil {
			return nil, err
		}
	}

	end := r.endp(pat)

	req, err := http.NewRequest(met, end, bytes.NewBuffer(b))
	if err != nil {
		return nil, tracer.Mask(err)
	}

	if dat != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	if r.pri {
		now := time.Now().UTC().Format("2006-01-02T15:04:05.000Z")

		req.Header.Set("DYDX-SIGNATURE", r.sign(pat, met, now, b))
		req.Header.Set("DYDX-API-KEY", r.sec.ApiKey)
		req.Header.Set("DYDX-TIMESTAMP", now)
		req.Header.Set("DYDX-PASSPHRASE", r.sec.ApiPas)
	}

	res, err := r.cli.Do(req)
	if err != nil {
		return nil, tracer.Mask(err)
	}
	defer res.Body.Close()

	b, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, tracer.Mask(err)
	}

	var lis errors.Errors
	err = json.Unmarshal(b, &lis)
	if err != nil {
		return nil, tracer.Mask(err)
	}

	if len(lis.Errors) != 0 {
		return nil, tracer.Mask(fmt.Errorf(lis.Errors[0].Msg))
	}

	return b, nil
}
