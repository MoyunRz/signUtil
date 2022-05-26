package util

import (
	"context"
	"fmt"
	"os/exec"
	"testing"
)

// 接收子协程的数据,协程之间用chan通信
type result struct {
	output []byte
	err    error
}

func Test_Key(t *testing.T) {
	// 执行一个cmd，让他在一个携程里面执行2s，
	/*
		1. WithCancel()函数接受一个 Context 并返回其子Context和取消函数cancel
		2. 新创建协程中传入子Context做参数，且需监控子Context的Done通道，若收到消息，则退出
		3. 需要新协程结束时，在外面调用 cancel 函数，即会往子Context的Done通道发送消息
		4. 注意：当 父Context的 Done() 关闭的时候，子 ctx 的 Done() 也会被关闭
	*/
	var ctx context.Context
	ctx, _ = context.WithCancel(context.TODO())
	// 生成命令
	cmd := exec.CommandContext(ctx, "bash", "-c", "../dimension-client transfer  --node-address=http://104.219.215.204:7782/rpc --chain-name=dimension-test --secret-key=../keygens/transfer/01863bb70d67edd2a3060f835f3f8807d7176834a5769c29600ecc3d09cfdc9c9b.pem --target-account=01863bb70d67edd2a3060f835f3f8807d7176834a5769c29600ecc3d09cfdc9c9b --amount=100 --payment-amount=100 --gas-price=500 --transfer-id=4")
	// 执行命令cmd.CombinedOutput(),且捕获输出
	output, err := cmd.CombinedOutput()
	fmt.Println(err)
	fmt.Println(string(output))
}
