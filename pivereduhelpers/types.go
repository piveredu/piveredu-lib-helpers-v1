package pivereduhelpers

import (
	"connectrpc.com/connect"
)

type ContextHelper interface {
	GetTenant(connect.AnyRequest) (string, error)
	GetAccessToken(connect.AnyRequest) (string, error)
}
