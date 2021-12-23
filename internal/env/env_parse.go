package env

import (
	"context"

	"github.com/sethvargo/go-envconfig"
)

var ctx = context.Background()

func Parse(config *Config) error {
	return envconfig.Process(ctx, config)
}
