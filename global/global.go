package global

import (
	"openapi/config"

	"github.com/redis/go-redis/v9"
)

var (
	// Config global config
	Config config.Config
	// Rdb redis client
	Rdb *redis.Client
)
