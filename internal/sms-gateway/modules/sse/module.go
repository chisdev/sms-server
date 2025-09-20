package sse

import (
	"context"

	"go.uber.org/fx"
	"go.uber.org/zap"
)

var Module = fx.Module(
	"sse",
	fx.Decorate(func(log *zap.Logger) *zap.Logger {
		return log.Named("sse")
	}),
	fx.Provide(
		newMetrics,
		fx.Private,
	),
	fx.Provide(
		NewService,
	),
	fx.Invoke(func(lc fx.Lifecycle, svc *Service) {
		lc.Append(fx.Hook{
			OnStop: func(ctx context.Context) error {
				return svc.Close(ctx)
			},
		})
	}),
)
