package main

import (
	"TTMS/configs/consts"
	"TTMS/internal/web/rpc"
	"TTMS/pkg/gmail"
	"TTMS/pkg/jwt"
)

func main() {
	rpc.InitRPC()
	jwt.InitRedis()
	gmail.New()
	r := InitRouter()
	err := r.Run("127.0.0.1:" + consts.WebServerPort)
	if err != nil {
		panic(err)
	}
}
