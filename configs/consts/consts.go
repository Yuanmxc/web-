package consts

import "time"

const (
	EtcdAddress     = "127.0.0.1:2379"
	NatsAddress     = "nats://127.0.0.1:4222"
	MySQLDefaultDSN = "TTMS:TTMS@tcp(localhost:3307)/TTMS?charset=utf8mb4&parseTime=True"
	RedisAddress    = "localhost:6378"
	RedisPassword   = "redis"
	RedisDB         = 0
	RedisTicketDB   = 1
	WebServerPort   = "8080"

	UserServiceName   = "userSvr"
	StudioServiceName = "studioSvr"
	PlayServiceName   = "playSvr"
	TicketServiceName = "ticketSvr"
	OrderServiceName  = "orderSvr"

	TicketCacheTime  = time.Hour * 24  //票缓存时间
	OrderDelayTime   = time.Minute * 3 //订单超时时间
	RedisLockTimeOut = time.Second * 5 //redis锁超时时间

	RPCTimeout     = 10 * time.Second
	ConnectTimeout = 5 * time.Second

	LimitRate = 2000.0 //令牌桶限流速率
	LimitCap  = 3000   //令牌桶容量

	JWTSecret = "kangning"
	// JWTOverTime 超长时间测试
	JWTOverTime = time.Hour * 72000
)
