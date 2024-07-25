package driver

import (
	"context"

	"github.com/roadrunner-server/api/v4/plugins/v4/jobs"
	"go.uber.org/zap"
)

var _ jobs.Driver = (*Driver)(nil)

type Configurer interface {
	// UnmarshalKey takes a single key and unmarshal it into a Struct.
	UnmarshalKey(name string, out any) error
	// Has checks if a config section exists.
	Has(name string) bool
}

type Driver struct {
	queue jobs.Queue
}

func FromConfig(configKey string, log *zap.Logger, cfg Configurer, pipeline jobs.Pipeline, pq jobs.Queue) (*Driver, error) { //nolint:revive
	return &Driver{
		queue: pq,
	}, nil
}

// FromPipeline initializes consumer from pipeline
func FromPipeline(pipeline jobs.Pipeline, log *zap.Logger, cfg Configurer, pq jobs.Queue) (*Driver, error) { //nolint:revive
	return &Driver{
		queue: pq,
	}, nil
}

func (d *Driver) Push(_ context.Context, job jobs.Message) error {
	item := fromJob(job)
	d.queue.Insert(item)
	return nil
}

func (d *Driver) Run(_ context.Context, _ jobs.Pipeline) error {
	return nil
}

func (d *Driver) State(_ context.Context) (*jobs.State, error) {
	return nil, nil
}

// Pause ctx + pipeline name
func (d *Driver) Pause(_ context.Context, _ string) error {
	return nil
}

// Resume ctx + pipeline name
func (d *Driver) Resume(_ context.Context, _ string) error {
	return nil
}

func (d *Driver) Stop(_ context.Context) error {
	return nil
}
