package main

import (
	"context"
	"fmt"

	"github.com/cloudwego/kitex/client"
	"github.com/showyquasar88/go-mall/demo/demo_thrift/kitex_gen/api"
	"github.com/showyquasar88/go-mall/demo/demo_thrift/kitex_gen/api/echo"
)

func main() {
	cli, err := echo.NewClient("demo_thrift", client.WithHostPorts("localhost:8888"))
	if err != nil {
		panic(err)
	}
	r, err := cli.Echo(context.Background(), &api.Request{
		Message: "你好",
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%v", r)
}