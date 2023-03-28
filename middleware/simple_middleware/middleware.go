package simple_middleware //nolint:revive,stylecheck

import (
	"net/http"

	"go.uber.org/zap"
)

const name = "sample-middleware"

type Configurer interface {
	// UnmarshalKey takes a single key and unmarshal it into a Struct.
	UnmarshalKey(name string, out any) error
	// Has checks if config section exists.
	Has(name string) bool
}

type Logger interface {
	NamedLogger(name string) *zap.Logger
}

type Plugin struct {
	log *zap.Logger
	cfg Configurer
}

// Init .. Logger and Configurer interfaces represents logger and configurer plugins.
// They would be provided automatically when registered in the RRs container.
func (p *Plugin) Init(logger Logger, cfg Configurer) error {
	// construct a named logger for the middleware
	p.log = logger.NamedLogger(name)
	p.cfg = cfg
	return nil
}

// Middleware is a special interface which after registration, RR will find automatically and add it to the http chain
func (p *Plugin) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		p.log.Info("i'm here")
		next.ServeHTTP(w, r)
	})
}

func (p *Plugin) Name() string {
	return name
}
