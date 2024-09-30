package models

import (
	"context"
	"strings"

	"github.com/redis/go-redis/v9"
)

type Savne struct {
	Password string
	Ip       string
	Key      string
}

func (s *Savne) ConnecttoRedis() (rdbs map[string]*redis.Client, err error) {
	rdbs = make(map[string]*redis.Client)
	ips := strings.Split(s.Ip, ",")
	for _, v := range ips {
		rdb := redis.NewClient(&redis.Options{
			Addr:     v,
			Password: s.Password, // no password set
			DB:       0,          // use default DB
		})
		err = rdb.Ping(context.Background()).Err()
		if err != nil {
			return nil, err
		}
		rdbs[v] = rdb
	}

	return rdbs, nil
}
