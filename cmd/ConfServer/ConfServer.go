package TinkConfigParser

import "io"

type ServerConfiguration struct {
	Port    int32
	Host    string
	Timeout int64
}

type IConfigParser interface {
	ParseConfig(r io.Reader) (ServerConfiguration, error)
}
