package config

import (
	"flag"
	"fmt"
	"redisscankey/models"
	"time"

	"github.com/redis/go-redis/v9"
	log "github.com/sirupsen/logrus"
)

var Savnes models.Savne
var RDB map[string]*redis.Client
var msg string

func init() {
	initlog()
	initflag()
	initredis()
}

func initlog() {
	log.SetFormatter(&log.JSONFormatter{
		TimestampFormat: time.DateTime,
	})
}
func initflag() {
	log.Info("初始化脚本")
	flag.StringVar(&Savnes.Ip, "ip", "127.0.0.1:6307", "当前的redis 地址（如果是集群，用逗号分隔多个地址）")
	flag.StringVar(&Savnes.Password, "password", "password", "当前的redis 密码")
	flag.StringVar(&Savnes.Key, "key", "key*", "当前的redis key")
	flag.Parse()
	log.Info("初始完成")
}

func initredis() {
	log.Info("开始进行连接redis")
	rdb, err := Savnes.ConnecttoRedis()
	if err != nil {
		msg = fmt.Sprintf("redis连接失败:%s", err)
		log.Error(msg)
		panic(err)
	}
	RDB = rdb
	log.Info("redis连接成功")
}
