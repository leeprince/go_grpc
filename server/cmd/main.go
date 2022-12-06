package main

import (
    "flag"
    "fmt"
    "github.com/leeprince/grpc_protocol/grpc_go/helloctl"
    "google.golang.org/grpc"
    "leeprince/go_grpc_server/internal/api"
    "log"
    "net"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022/11/29 00:49
 * @Desc:   RPC 服务端
 */

var (
    port = flag.Int("port", 8000, "server port")
)

func main() {
    flag.Parse()
    
    listenner, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
    if err != nil {
        log.Fatalf("net.Listen err:", err)
    }
    
    // 实例化 HelloCtl 服务端
    gprcServer := grpc.NewServer()
    helloctl.RegisterHelloCtlServer(gprcServer, api.NewHelloCtl())
    
    log.Printf("server listening at %v", listenner.Addr())
    if err := gprcServer.Serve(listenner); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}
