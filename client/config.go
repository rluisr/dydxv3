package client

import "github.com/rluisr/dydxv3/client/secret"

type Config struct {
	Sec secret.Secret
	Tes bool
}

func (c Config) Verify() {}
