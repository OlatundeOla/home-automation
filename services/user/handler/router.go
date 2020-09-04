// Code generated by jrpc. DO NOT EDIT.

package handler

import (
	context "context"
	http "net/http"

	taxi "github.com/jakewright/home-automation/libraries/go/taxi"
	def "github.com/jakewright/home-automation/services/user/def"
)

// taxiRouter is an interface implemented by taxi.Router
type taxiRouter interface {
	HandleFunc(method, path string, handler func(context.Context, taxi.Decoder) (interface{}, error))
}

type handler interface {
	GetUser(ctx context.Context, body *def.GetUserRequest) (*def.GetUserResponse, error)
	ListUsers(ctx context.Context, body *def.ListUsersRequest) (*def.ListUsersResponse, error)
}

// RegisterRoutes adds the service's routes to the router
func RegisterRoutes(r taxiRouter, h handler) {
	r.HandleFunc("GET", "/user", func(ctx context.Context, decode taxi.Decoder) (interface{}, error) {
		body := &def.GetUserRequest{}
		if err := decode(body); err != nil {
			return nil, err
		}

		if err := body.Validate(); err != nil {
			return nil, err
		}

		return h.GetUser(ctx, body)
	})

	r.HandleFunc("GET", "/users", func(ctx context.Context, decode taxi.Decoder) (interface{}, error) {
		body := &def.ListUsersRequest{}
		if err := decode(body); err != nil {
			return nil, err
		}

		if err := body.Validate(); err != nil {
			return nil, err
		}

		return h.ListUsers(ctx, body)
	})

}

// newHandler returns a handler that serves requests for
// this service. This is not exported as it is only used by
// tests. A service's main() function should create its own
// router (typically via bootstrap) and then use RegisterRoutes().
func newHandler(h handler) http.Handler {
	r := taxi.NewRouter()
	RegisterRoutes(r, h)
	return r
}