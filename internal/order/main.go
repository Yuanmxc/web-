package main

import (
	"TTMS/configs/consts"
	"TTMS/internal/order/dao"
	"TTMS/internal/order/mw"
	"TTMS/internal/order/service"
	order "TTMS/kitex_gen/order/orderservice"
	"net"
	"time"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
)

func main() {
	r, err := etcd.NewEtcdRegistry([]string{consts.EtcdAddress}) // r should not be reused.
	if err != nil {
		panic(err)
	}
	addr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:10005")
	if err != nil {
		panic(err)
	}
	svr := order.NewServer(new(OrderServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: consts.OrderServiceName}), // server name
		server.WithReadWriteTimeout(10*time.Second),
		server.WithServiceAddr(addr),                                         // address
		server.WithLimit(&limit.Option{MaxConnections: 10000, MaxQPS: 5000}), // limit
		//server.WithMuxTransport(),                                            // Multiplex，win不支持
		//server.WithSuite(trace.NewDefaultServerSuite()),                     // tracer
		server.WithRegistry(r), // registry
	)
	mw.LoadLocation()
	dao.Init()
	mw.InitRedis()
	service.InitPlayRPC()
	go mw.InitNats()
	err = svr.Run()
	if err != nil {
		klog.Fatal(err)
	}
}
