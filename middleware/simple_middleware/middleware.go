package simple_middleware //nolint:revive,stylecheck

import (
	"net/http"

	"go.uber.org/zap"
)

type Plugin struct {
	log *zap.Logger
}

func (p *Plugin) Init(log *zap.Logger) error {
	p.log = &zap.Logger{}
	*p.log = *log
	return nil
}

// Middleware is a special interface which after registration, RR will find automatically and add it to the http chain
func (p *Plugin) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		p.log.Info("i'm here")
		next.ServeHTTP(w, r)
	})
}
