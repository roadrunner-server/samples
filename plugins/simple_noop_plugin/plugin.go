// Package simple_noop_plugin demonstrates a minimal no-op RoadRunner plugin.
package simple_noop_plugin //nolint:revive

import (
	"context"

	"github.com/roadrunner-server/endure/v2/dep"
	"go.uber.org/zap"
)

const name = "sample-plugin"

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
	myF F
}

// Init .. Logger and Configurer interfaces represents logger and configurer plugins.
// They would be provided automatically when registered in the RRs container.
func (p *Plugin) Init(logger Logger, cfg Configurer) error {
	// construct a named logger for the middleware
	p.log = logger.NamedLogger(name)
	p.cfg = cfg
	return nil
}

func (p *Plugin) Serve() chan error {
	errCh := make(chan error, 1)

	/*
		your logic here
	*/

	return errCh
}

func (p *Plugin) FooBar() string {
	return "foobar"
}

func (p *Plugin) Weight() uint {
	return 100
}

type F interface {
	FooBar() string
}

// Collects all plugins which implement Name + RPCer interfaces
func (p *Plugin) Collects() []*dep.In {
	return []*dep.In{
		dep.Fits(func(plugin any) {
			pp := plugin.(F)
			p.myF = pp
		}, (*F)(nil)),
	}
}

// Stop would be called on the RR stop.
func (p *Plugin) Stop(_ context.Context) error {
	return nil
}

// Name interface should be implemented to provide a user-friendly plugin name.
func (p *Plugin) Name() string {
	return name
}
