package util

import (
	"context"
	"fmt"
	"github.com/bsm/redislock"
	"github.com/go-redis/redis/v8"
	"testing"
	"time"
)

var ctx = context.Background()
var rdb *redis.Client

func init() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		DB:       0,
		Password: "123456",
	})
}

func Test_redis1(t *testing.T) {
	stat := rdb.Set(ctx, "xxx", "aaa", -1)
	fmt.Println(stat)
}

func Test_redis2(t *testing.T) {
	val, err := rdb.Get(ctx, "xxx").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(val)
}

func Test_redis3(t *testing.T) {
	iter := rdb.Scan(ctx, 0, "x*", 0).Iterator()
	for iter.Next(ctx) {
		fmt.Println("keys", iter.Val())
	}
	if err := iter.Err(); err != nil {
		panic(err)
	}
}

func Test_redis4(t *testing.T) {
	// 获取一个不存在的key时，返回 redis.Nil err
	val, err := rdb.Get(ctx, "xxxxx").Result()
	if err == redis.Nil {
		panic(err)
	}
	fmt.Println(val)
}

func Test_redis5(t *testing.T) {
	// 数字操作 - 自增
	val, err := rdb.Incr(ctx, "num").Result()
	fmt.Println("err", err)
	fmt.Println("val", val)

	// 自减
	val2, err := rdb.Decr(ctx, "num2").Result()
	fmt.Println("err", err)
	fmt.Println("val2", val2)
}

// 分布式锁
func TestRedisLockImpl(t *testing.T) {
	locker := redislock.New(rdb)
	// 获取锁
	lock, err := locker.Obtain(ctx, "my-key", 100*time.Millisecond, nil)
	if err != nil {
		fmt.Println("Could not obtain lock!")
	}
	fmt.Println("get lock success, do work.....")
	// 释放锁
	lock.Release(ctx)
	time.Sleep(50 * time.Millisecond)
}
