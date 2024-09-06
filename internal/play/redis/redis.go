package redis

import (
	"TTMS/configs/consts"
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
	"strconv"
)

var redisClient *redis.Client

func Init() {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     consts.RedisAddress,
		Password: consts.RedisPassword,
		DB:       consts.RedisTicketDB,
	})
}

func UpdatePlayPrice(ctx context.Context, schList []int64, price int) {
	m := make(map[string]string, len(schList))
	for _, sch := range schList {
		m[fmt.Sprintf("%d;price", sch)] = strconv.Itoa(price)
	}
	log.Println(m)
	redisClient.MSet(ctx, m)
}
