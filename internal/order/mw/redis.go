package mw

import (
	"TTMS/configs/consts"
	"TTMS/internal/order/dao"
	"TTMS/kitex_gen/order"
	"context"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
	"strconv"
	"strings"
	"time"
)

/*
+---------------------------------------------------+
|					delayQueue->member				|
+---------------------------------------------------+
|	UserId	|	ScheduleId	|	SeatRow	|	SeatCol	|
+---------------------------------------------------+
*/
var client *redis.Client
var delayQueue = "delay_queue"
var targetQueue = "target_queue"
var Loc *time.Location

func InitRedis() {
	client = redis.NewClient(&redis.Options{
		Addr:     consts.RedisAddress,
		Password: consts.RedisPassword,
		DB:       consts.RedisTicketDB,
	})
	ctx := context.Background()
	go toTargetQueue(ctx)
	go eventLoop(ctx)
}
func LoadLocation() {
	var err error
	Loc, err = time.LoadLocation("Asia/Shanghai")
	if err != nil {
		log.Println("Error loading location:", err)
		return
	}
}

// ToDelayQueue 将任务添加到延迟队列
func ToDelayQueue(ctx context.Context, orderInfo string, timeUnix float64) {
	err := client.ZAdd(ctx, delayQueue, redis.Z{Member: orderInfo, Score: timeUnix}).Err()
	log.Println("ToDelayQueue time = ", time.Unix(int64(timeUnix), 0).Format("2006-01-02 15:04:05"))
	if err != nil {
		log.Println("ToDelayQueue ", err)
	}
}

// RemoveFromDelayQueue 按时支付订单，从延时队列中取出该订单
func RemoveFromDelayQueue(ctx context.Context, req *order.CommitOrderRequest) error {
	orderInfo := fmt.Sprintf("%d;%d;%d;%d", req.UserId, req.ScheduleId, req.SeatRow, req.SeatCol)
	log.Println("orderInfo = ", orderInfo)
	count, err := client.ZRem(ctx, delayQueue, orderInfo).Result()
	log.Println("count = ", count, " err = ", err)
	if err != nil {
		return err
	}
	if count == 0 {
		return errors.New("订单已经过期")
	}
	//更改订单状态为已支付
	err = dao.UpdateOrder(req.UserId, req.ScheduleId, req.SeatRow, req.SeatCol, 1, time.Now().Format("2006-01-02 15:04:05"))
	if err != nil {
		log.Println(err)
		return err
	}
	//向ticket服务提交，更改票状态为已付款
	_, err = JS.Publish("stream.ticket.commit", []byte(fmt.Sprintf("%d;%d;%d", req.ScheduleId, req.SeatRow, req.SeatCol)))
	if err != nil {
		log.Panicln("UpdateOrder err = ", err)
	}
	return nil
}

// toTargetQueue 处理延迟队列取出来的订单
func toTargetQueue(ctx context.Context) {
	for {
		// 从延迟队列中取出最小的时间戳
		result, err := client.ZRangeByScoreWithScores(ctx, delayQueue, &redis.ZRangeBy{
			Min:    "-inf",
			Max:    fmt.Sprintf("%d", time.Now().Unix()),
			Offset: 0,
			Count:  100,
		}).Result()
		if err != nil {
			// 处理错误
			panic(err)
		}

		if len(result) == 0 {
			// 延迟队列中没有数据，等待一段时间后再次查询
			time.Sleep(time.Second)
			continue
		}

		for _, z := range result {
			// 从延迟队列中移除该数据
			_, err := client.ZRem(ctx, delayQueue, z.Member).Result()
			if err != nil {
				// 处理错误
				panic(err)
			}

			// 将数据从延迟队列转移到目标队列
			fmt.Println(z.Member, z.Score)
			_, err = client.LPush(ctx, targetQueue, z.Member).Result()
			if err != nil {
				// 处理错误
				panic(err)
			}
		}
	}
}

// 循环处理targetQueue中的数据 (这些都是超时未支付的订单，需要从数据库中删除order实体，重新放出票)
func eventLoop(ctx context.Context) {
	for {
		results, err := client.BLPop(ctx, time.Second*5, targetQueue).Result()
		if err == redis.Nil {
			continue
		}
		log.Println("now = ", time.Now().Format("2006-01-02 15:04:05"))
		log.Println("results = ", results)
		if len(results) == 0 {
			continue
		}
		results = results[1:]
		log.Println("results = ", results)
		for _, result := range results {
			fmt.Println("result = ", result)
			data := strings.Split(result, ";")
			fmt.Println("data = ", data)
			d0, _ := strconv.Atoi(data[0])
			d1, _ := strconv.Atoi(data[1])
			d2, _ := strconv.Atoi(data[2])
			d3, _ := strconv.Atoi(data[3])
			//向ticket服务提交，更改票状态为 待售
			_, err = JS.Publish("stream.ticket.timeout", []byte(fmt.Sprintf("%d;%d;%d", d1, d2, d3)))
			if err != nil {
				log.Panicln(err)
			}
			err = dao.DeleteOrder(ctx, d0, d1, d2, d3)
			if err != nil {
				log.Println(err)
			}
		}

	}
}
