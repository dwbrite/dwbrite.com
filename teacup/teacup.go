package teacup

import (
	"regexp"
		)

type teacupServer struct {
	port uint16

	fileWhitelist *regexp.Regexp
	dirBlacklist *regexp.Regexp
}

var DefaultTeacup = &teacupServer{41234, nil, nil}

func New(port uint16, cert string, key string) *teacupServer {
	return &teacupServer{port, nil, nil}
}

