package util

import (
	"context"
	"encoding/json"
	"log"
	"os/exec"
)

func BashByCommand(cmdStr string, res interface{}) error {
	var ctx context.Context
	ctx, _ = context.WithCancel(context.TODO())
	// 生成命令
	cmd := exec.CommandContext(ctx, "bash", "-c", cmdStr)
	log.Printf("执行命令 :%s", cmdStr)
	// 执行命令cmd.CombinedOutput(),且捕获输出
	output, err := cmd.CombinedOutput()
	log.Printf("执行命令 output:%s", string(output))
	if err != nil {
		log.Printf("执行命令 错误:%v", err)
		return err
	}
	if res != nil {
		err = json.Unmarshal(output, res)
		if err != nil {
			return err
		}
	}
	return err
}
func CallByCommand(cmdStr string, res interface{}) error {
	var ctx context.Context
	ctx, _ = context.WithCancel(context.TODO())
	// 生成命令
	cmd := exec.CommandContext(ctx, "bash", cmdStr)
	log.Printf("执行命令 :%s", cmdStr)
	// 执行命令cmd.CombinedOutput(),且捕获输出
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("执行命令 错误:%s", err.Error())
		return err
	}
	log.Printf("执行命令 output:%s", string(output))
	if res != nil {
		err = json.Unmarshal(output, res)
		if err != nil {
			return err
		}
	}
	return err
}
