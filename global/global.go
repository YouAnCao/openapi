package global

import (
	"openapi/config"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var (
	// Config global config
	Config config.Config
	// Rdb redis client
	Rdb *redis.Client
	Db  *gorm.DB
)
