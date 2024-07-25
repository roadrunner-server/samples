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
}

func FromConfig(configKey string, log *zap.Logger, cfg Configurer, pipeline jobs.Pipeline, pq jobs.Queue) (*Driver, error) {
	return &Driver{}, nil
}

// FromPipeline initializes consumer from pipeline
func FromPipeline(pipeline jobs.Pipeline, log *zap.Logger, cfg Configurer, pq jobs.Queue) (*Driver, error) {
	return &Driver{}, nil
}

func (d *Driver) Push(ctx context.Context, job jobs.Message) error {
	return nil
}

func (d *Driver) Run(ctx context.Context, p jobs.Pipeline) error {
	return nil
}

func (d *Driver) State(ctx context.Context) (*jobs.State, error) {
	return nil, nil
}

func (d *Driver) Pause(ctx context.Context, p string) error {
	return nil
}

func (d *Driver) Resume(ctx context.Context, p string) error {
	return nil
}

func (d *Driver) Stop(ctx context.Context) error {
	return nil
}
