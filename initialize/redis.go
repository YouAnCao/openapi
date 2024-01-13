package initialize

import (
	"fmt"
	"openapi/global"

	"github.com/redis/go-redis/v9"
)

// Redis init
func Redis() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", global.Config.Redis.Host, global.Config.Redis.Port),
		Password: "",
		DB:       0,
	})
	global.Rdb = rdb
}
