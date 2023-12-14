/**
 * Created by lock
 * Date: 2019-08-12
 * Time: 14:18
 */
package tools

import (
	"fmt"
	"github.com/go-redis/redis"
	"sync"
	"time"
)

var RedisClientMap = map[string]*redis.Client{}
var syncLock sync.Mutex

// RedisOption redis 连接配置信息
type RedisOption struct {
	Address  string
	Password string
	Db       int
}

func GetRedisInstance(redisOpt RedisOption) *redis.Client {
	address := redisOpt.Address
	db := redisOpt.Db
	password := redisOpt.Password
	addr := fmt.Sprintf("%s", address)
	// 获取 redis 对象， 若存在则直接返回，不存在则新建一个，这个通过 map 实现单例模式
	if redisCli, ok := RedisClientMap[addr]; ok {
		return redisCli
	}
	// 这里源码中是在获取的时候就开始加锁来，我觉得不太合理，所以修改了此处代码
	// 将加锁延时到获取之后创建之前
	// 加写锁
	syncLock.Lock()
	// 创建 redis 对象
	client := redis.NewClient(&redis.Options{
		Addr:       addr,
		Password:   password,
		DB:         db,
		MaxConnAge: 20 * time.Second,
	})
	// 将 redis 对象放入map
	RedisClientMap[addr] = client
	// 解锁
	syncLock.Unlock()
	// 返回 redis 实例
	return RedisClientMap[addr]
}
