package test

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"testing"
)

var ctx = context.Background()

func TestRedis(t *testing.T) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	err := rdb.Set(ctx, "1001", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	get := rdb.Get(ctx, "1001")
	fmt.Println(get)

}
