package simple_noop_plugin //nolint:revive,stylecheck

import (
	"go.uber.org/zap"
)

type Plugin struct{}

// Init method will be called first for the plugin
func (p *Plugin) Init(log *zap.Logger) error {
	log.Info("i'm started")
	return nil
}
