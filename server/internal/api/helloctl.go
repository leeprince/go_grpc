package api

import (
    "context"
    "fmt"
    "github.com/leeprince/grpc_protocol/grpc_go/helloctl"
    "time"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022/11/29 18:37
 * @Desc:
 */

// --- 实现 helloctl.HelloCtlServer 接口
type HelloCtl struct {
    helloctl.UnimplementedHelloCtlServer
}

func NewHelloCtl() *HelloCtl {
    return new(HelloCtl)
}
func (ctl *HelloCtl) SayHello(ctx context.Context, req *helloctl.SayHelloReq) (*helloctl.SayHelloResp, error) {
    fmt.Println("helloctl.SayHelloReq", req)
    
    resp := &helloctl.SayHelloResp{
        Code:    0,
        Message: "success",
        LogId:   "prince-test",
        Data: &helloctl.SayHelloRespData{
            AccessTime: time.Now().Unix(),
        },
    }
    
    return resp, nil
}
