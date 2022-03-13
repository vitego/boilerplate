package middlewares

import (
	"github.com/vitego/router"
	"github.com/vitego/router/mw"
)

type Handler struct{}

// Defaults contains middleware that would be execute on every request
func (Handler) Defaults() []router.Middleware {
	return []router.Middleware{
		mw.RequestMiddleware{},
	}
}

// Middleware contains usable middleware in your controllers
func (Handler) Middleware() map[string]router.Middleware {
	return map[string]router.Middleware{
		"toto": mw.RequestMiddleware{},
	}
}
