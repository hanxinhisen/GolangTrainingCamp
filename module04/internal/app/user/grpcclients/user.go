package grpcclients

import (
	userproto "blog-users/api/protos/user"
	"blog-users/internal/pkg/etcdservice"
	"blog-users/internal/pkg/transports/grpc"
	"fmt"
	"log"

	"github.com/pkg/errors"
	"github.com/spf13/viper"
	grpc2 "google.golang.org/grpc"
	"google.golang.org/grpc/balancer/roundrobin"
	"google.golang.org/grpc/resolver"
)

// NewUserClient 初始化 user rpc client
func NewUserClient(client *grpc.Client, v *viper.Viper) (userproto.UserClient, error) {
	o := new(clientTarget)
	if err := v.UnmarshalKey("client.target", o); err != nil {
		return nil, errors.Wrap(err, "获取client.target配置失败")
	}
	// 兼容现阶段生产环境无etcd服务的问题 根据配置文件执行是否etcd调用
	//conn, err := client.Dial(o.User, grpc.WithGrpcDialOptions(grpc2.WithInsecure()))
	conn, err := client.Dial(o.User)
	fmt.Println("conn:", conn)
	if o.User == "" {
		log.Println("执行etcd调用")
		//注册etcd解析器
		r := etcdservice.NewResolver(o.EtcdAddr)
		resolver.Register(r)

		// 客户端连接服务器
		conn, err = grpc2.Dial(r.Scheme()+"://"+o.Caller+"/"+o.Callee, grpc2.WithDefaultServiceConfig(fmt.Sprintf(`{"LoadBalancingPolicy": "%s"}`, roundrobin.Name)), grpc2.WithInsecure())
	}

	if err != nil {
		log.Println("连接服务器失败", err)
		return nil, errors.Wrap(err, "notify client dial error")
	}

	return userproto.NewUserClient(conn), nil
}
