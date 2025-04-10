package pivereduhelpers

import (
	"connectrpc.com/connect"
	"errors"
	"strings"
)

const (
	XTenantKey = "x-tenant-id"
)

type contextHelper struct {
}

func (contextHelper *contextHelper) GetAccessToken(request connect.AnyRequest) (string, error) {
	// Look for the authorization header.
	authHeader := request.Header().Get("Authorization")
	if authHeader == "" {
		return "", connect.NewError(connect.CodeUnauthenticated, errors.New("no authorization header"))
	}

	// The authorization header should be in the form "Bearer <token>".
	parts := strings.SplitN(authHeader, " ", 2)
	if len(parts) != 2 || !strings.EqualFold(parts[0], "Bearer") {
		return "", connect.NewError(connect.CodeUnauthenticated, errors.New("invalid authorization header"))
	}

	return parts[1], nil
}

func (contextHelper *contextHelper) GetTenant(req connect.AnyRequest) (string, error) {
	tenantId := req.Header().Get(XTenantKey)
	if tenantId == "" {
		return "", connect.NewError(connect.CodeUnauthenticated, errors.New("no tenant id passed in request"))
	}

	return tenantId, nil
}

func NewContextHelper() ContextHelper {
	return &contextHelper{}
}
