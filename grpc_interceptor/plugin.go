// Package interceptor demonstrates a custom RoadRunner gRPC interceptor plugin.
package interceptor

import (
	"context"
	"time"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

const name = "sample-grpc-interceptor"

type Logger interface {
	NamedLogger(name string) *zap.Logger
}

type Plugin struct {
	log *zap.Logger
}

func (p *Plugin) Init(logger Logger) error {
	p.log = logger.NamedLogger(name)
	return nil
}

func (p *Plugin) UnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
		started := time.Now()

		resp, err := handler(ctx, req)
		if err != nil {
			st, _ := status.FromError(err)
			p.log.Warn(
				"grpc request failed",
				zap.String("method", info.FullMethod),
				zap.String("status", st.Code().String()),
				zap.Int64("elapsed_ms", time.Since(started).Milliseconds()),
				zap.Error(err),
			)

			return resp, err
		}

		p.log.Info(
			"grpc request completed",
			zap.String("method", info.FullMethod),
			zap.Int64("elapsed_ms", time.Since(started).Milliseconds()),
		)

		return resp, nil
	}
}

func (p *Plugin) Name() string {
	return name
}
