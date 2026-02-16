// Package jobs_driver demonstrates how to register a custom RoadRunner jobs driver.
package jobs_driver //nolint:revive

import (
	"github.com/roadrunner-server/api/v4/plugins/v4/jobs"
	"github.com/roadrunner-server/errors"
	"github.com/roadrunner-server/samples/plugins/jobs_driver/driver"
	"go.uber.org/zap"
)

const pluginName string = "my_awesome_driver"

type Configurer interface {
	// UnmarshalKey takes a single key and unmarshal it into a Struct.
	UnmarshalKey(name string, out any) error
	// Has checks if a config section exists.
	Has(name string) bool
}

type Logger interface {
	NamedLogger(name string) *zap.Logger
}

type Plugin struct {
	log *zap.Logger
	cfg Configurer
}

func (p *Plugin) Init(log Logger, cfg Configurer) error {
	if !cfg.Has(pluginName) {
		return errors.E(errors.Disabled)
	}

	p.log = log.NamedLogger(pluginName)
	p.cfg = cfg
	return nil
}

func (p *Plugin) Name() string {
	return pluginName
}

// DriverFromConfig constructs kafka driver from the .rr.yaml configuration
func (p *Plugin) DriverFromConfig(configKey string, pq jobs.Queue, pipeline jobs.Pipeline) (jobs.Driver, error) {
	return driver.FromConfig(configKey, p.log, p.cfg, pipeline, pq)
}

// DriverFromPipeline constructs kafka driver from pipeline
func (p *Plugin) DriverFromPipeline(pipe jobs.Pipeline, pq jobs.Queue) (jobs.Driver, error) {
	return driver.FromPipeline(pipe, p.log, p.cfg, pq)
}
