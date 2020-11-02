package initialize

import (
	"github.com/go-redis/redis"
	"tracking/global"
)

func Redis() {
	client := redis.NewClient(&redis.Options{
		Addr:     global.GVA_CONFIG.Redis.Addr,
		Password: global.GVA_CONFIG.Redis.Password,
		DB:       global.GVA_CONFIG.Redis.DB,
	})

	if pong, err := client.Ping().Result(); err != nil {
		global.GVA_LOG.Error()
	} else {
		global.GVA_LOG.Info("redis connect ping response:", pong)
		global.GVA_REDIS = client
	}
}
