package controllers

import (
	"context"

	"github.com/redis/go-redis/v9"
)

// Redis_key这个是进行模糊匹配key的总值
func Redis_key(rdb *redis.Client, key_prefix string) (int64, error) {
	var cursor uint64
	var n int64
	for {
		var keys []string
		var err error
		keys, cursor, err = rdb.Scan(context.Background(), cursor, key_prefix, 10).Result()
		if err != nil {
			return 0, err
		}
		n += int64(len(keys))
		if cursor == 0 {
			break
		}
	}
	return n, nil

}
