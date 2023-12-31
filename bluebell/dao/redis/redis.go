package redis

import (
	"bluebell/settings"
	"fmt"

	"github.com/spf13/viper"

	"github.com/go-redis/redis"
)

// 声明一个全局的rdb变量
var rdb *redis.Client

// 初始化连接
func InitRedis(cfg *settings.RedisConfig) (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:         fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Password:     cfg.Password, // no password set
		DB:           cfg.DB,       // use default DB
		PoolSize:     cfg.PoolSize,
		MinIdleConns: cfg.MinIdleConns,
	})
	fmt.Println("打印....")
	fmt.Println(fmt.Sprintf("%s:%d", viper.GetString(cfg.Host), viper.GetInt("cfg.Port")))
	_, err = rdb.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}
func Close() {
	_ = rdb.Close()
}
