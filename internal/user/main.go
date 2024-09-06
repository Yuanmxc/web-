package main

import (
	"TTMS/configs/consts"
	"TTMS/internal/user/dao"
	user "TTMS/kitex_gen/user/userservice"
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
	addr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:10001")
	if err != nil {
		panic(err)
	}
	svr := user.NewServer(new(UserServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: consts.UserServiceName}), // server name
		//server.WithMiddleware(mw.CommonMiddleware),                                          // middleWare
		//server.WithMiddleware(mw.ServerMiddleware),
		server.WithServiceAddr(addr),                                         // address
		server.WithLimit(&limit.Option{MaxConnections: 10000, MaxQPS: 5000}), // limit
		server.WithReadWriteTimeout(10*time.Second),
		//server.WithMuxTransport(),                                            // 开启多路复用，win不支持
		//server.WithSuite(trace.NewDefaultServerSuite()),                     // tracer
		server.WithRegistry(r), // registry
	)
	dao.Init()
	err = svr.Run()
	if err != nil {
		klog.Fatal(err)
	}
}
