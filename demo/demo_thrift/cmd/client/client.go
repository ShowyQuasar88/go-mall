package main

import (
	"context"
	"fmt"
	"log"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/transmeta"
	"github.com/cloudwego/kitex/transport"
	consul "github.com/kitex-contrib/registry-consul"
	"github.com/showyquasar88/go-mall/demo/demo_thrift/kitex_gen/api"
	"github.com/showyquasar88/go-mall/demo/demo_thrift/kitex_gen/api/echo"
)

func main() {
	// dest service
	// cli, err := echo.NewClient("demo_thrift", client.WithHostPorts("localhost:8888"))

	// use consul resolver find with name
	r, err := consul.NewConsulResolver("127.0.0.1:8500")
	if err != nil {
		log.Fatal(err)
	}
	cli, err := echo.NewClient("demo_thrift", client.WithResolver(r),
		client.WithMetaHandler(transmeta.ClientTTHeaderHandler),
		client.WithTransportProtocol(transport.TTHeader),
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: "demo_thrift_client",
		}),
	)
	if err != nil {
		panic(err)
	}
	res, err := cli.Echo(context.Background(), &api.Request{
		Message: "你好",
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%v", res)
}
