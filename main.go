package main

import (
	"context"
	"fmt"
	"redisscankey/config"
	"redisscankey/controllers"
	"sync"

	"github.com/redis/go-redis/v9"
	log "github.com/sirupsen/logrus"
)

var dbsizes int64
var keys int64
var msg string
var wg sync.WaitGroup

func main() {

	log.Info("开始执行")
	// 开始进行获取redis 所有的key数量
	for k, v := range config.RDB {
		wg.Add(1)
		go func(k string, v *redis.Client) { // 替换 YourRedisClientType 为实际类型
			defer wg.Done()
			dbsize := v.DBSize(context.Background()).Val()
			msg = fmt.Sprintf("当前节点：%s 的key数：%d", k, dbsize)
			log.Info(msg)
			dbsizes += dbsize
			// 进行循环匹配key
			key, err := controllers.Redis_key(v, config.Savnes.Key)
			if err != nil {
				log.Error(err)
			} else {
				msg = fmt.Sprintf("当前节点：%s 模糊匹配key数：%d", k, key)
				log.Info(msg)
			}
			keys += key
		}(k, v) // 将 k 和 v 作为参数传入 goroutine
	}

	wg.Wait()
	fmt.Println("当前redis一共有", dbsizes)
	// 开始进行模糊匹配有多少key
	fmt.Println("模糊匹配key有", keys)
	log.Info("执行完成")
}
