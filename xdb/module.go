package xdb

import (
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		NewDB,
		DBHealthCheck,
		LoadConfig,
	),
)
