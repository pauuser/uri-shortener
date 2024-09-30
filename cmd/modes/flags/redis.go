package flags

import "github.com/redis/go-redis/v9"

type RedisFlags struct {
	Address  string `mapstructure:"address"`
	Database int    `mapstructure:"database"`
	Password string `mapstructure:"password"`
}

func (r *RedisFlags) InitRedis() (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     r.Address,
		Password: r.Password,
		DB:       r.Database,
	})

	return client, nil
}
