package main

import (
	"context"
	"fmt"
	"log"

	"github.com/cloudwego/kitex/client"
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
	cli, err := echo.NewClient("demo_thrift", client.WithResolver(r))
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