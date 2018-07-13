package conn

import (
	"fmt"
	"time"

	"github.com/garyburd/redigo/redis"
)

var redisClient *redis.Pool

type redisConf struct {
	host      string
	port      int64
	password  string
	maxIdle   int
	maxActive int
	db        int64
}

func init() {
	redisClient = newRedisPools()
}

func newRedisPools() *redis.Pool {
	config := getConf()
	return &redis.Pool{
		MaxIdle:     config.maxIdle,
		MaxActive:   config.maxActive,
		IdleTimeout: 30 * time.Second,
		Dial: func() (redis.Conn, error) {
			return setDialog(config)
		},
	}
}

func getConf() *redisConf {
	return &redisConf{host: "127.0.0.1", port: 6379, password: "mypassword",
		maxIdle: 10, maxActive: 10000, db: 2}
}

func setDialog(config *redisConf) (redis.Conn, error) {
	conn, err := redis.Dial("tcp", fmt.Sprintf("%s:%d", config.host, config.port))
	if err != nil {
		return nil, fmt.Errorf("init redis failure:%v", err)
	}
	if len(config.password) != 0 {
		if _, err := conn.Do("AUTH", config.password); err != nil {
			conn.Close()
			return nil, fmt.Errorf("init redis failure:%v", err)
		}
	}
	if _, err := conn.Do("SELECT", config.db); err != nil {
		conn.Close()
		return nil, fmt.Errorf("init redis failure:%v", err)
	}
	return conn, nil
}

// Get 获取redis的链接
func Get() redis.Conn {
	if redisClient == nil {
		panic("redis connection failed")
	}
	return redisClient.Get()
}
