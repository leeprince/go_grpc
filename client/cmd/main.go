package main

import (
    "context"
    "flag"
    "fmt"
    "github.com/leeprince/grpc_protocol/grpc_go/helloctl"
    "google.golang.org/grpc"
    "google.golang.org/grpc/credentials/insecure"
    "log"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022/11/29 14:07
 * @Desc:   RPC 客户端
 */

var (
    grpcServerHost = flag.String("grpc-server-host", "localhost:8000", "grpc server host")
)

func main() {
    flag.Parse()
    
    clientConn, err := grpc.Dial(*grpcServerHost, grpc.WithTransportCredentials(insecure.NewCredentials()))
    if err != nil {
        log.Fatalf("grpc.Dial err: %v", err)
    }
    defer clientConn.Close()
    
    // 实例化 helloctl 客户端
    helloCtlClient := helloctl.NewHelloCtlClient(clientConn)
    
    req := &helloctl.SayHelloReq{
        Name: "prince",
        Age:  18,
    }
    // 调用 helloCtl 的SayHello方法
    resp, err := helloCtlClient.SayHello(context.Background(), req)
    if err != nil {
        log.Fatalf("helloCtlClient.SayHello err: %v", err)
    }
    fmt.Printf("helloCtlClient.SayHello resp: %+v", resp)
}
