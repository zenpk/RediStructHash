package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
	"time"
)

type Video struct {
	Id         int64
	PlayUrl    string
	Title      string `redistructhash:"no"` // use "no" tag to prevent creating field
	CreateTime int64
}

func main() {
	var rdb *redis.Client
	var ctx = context.Background()
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	video := Video{
		Id:         1,
		PlayUrl:    "youtube.com",
		Title:      "my video",
		CreateTime: time.Now().Unix(),
	}
	if err := RedisStructHash(rdb, ctx, video, "key"); err != nil {
		log.Println(err)
	}
	fmt.Println(rdb.HGetAll(ctx, "key").Result())
}
