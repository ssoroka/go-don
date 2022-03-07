package don

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Empty struct{}

var DefaultEncoding = "text/plain"

type Middleware func(http.Handler) http.Handler

type API struct {
	*router
	NotFound         http.Handler
	MethodNotAllowed http.Handler
}

type Config struct {
	DefaultEncoding string
}

// New creates a new API instance.
func New(c *Config) *API {
	if c == nil {
		c = &Config{}
	}
	if c.DefaultEncoding == "" {
		c.DefaultEncoding = DefaultEncoding
	}
	return &API{
		router:           &router{config: c},
		NotFound:         E(ErrNotFound),
		MethodNotAllowed: E(ErrMethodNotAllowed),
	}
}

// Router creates a http.Handler for the API.
func (a *API) Router() http.Handler {
	rr := httprouter.New()
	a.addRoutes(rr)

	rr.NotFound = withConfig(a.NotFound, a.config)
	rr.MethodNotAllowed = withConfig(a.MethodNotAllowed, a.config)

	return rr
}
