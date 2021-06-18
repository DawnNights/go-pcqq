package main

import (
	"fmt"
	"pcqq/core"
)

func main() {
	pc := core.PCQQ{}
	pc.Init()
	pc.GetQrCode()

	var msg string
	var groupId int64 = 522245324	//填测试群号

	for {
		fmt.Print("请输入发送内容: ")
		fmt.Scanln(&msg)
		pc.SendGroupMsg(groupId,msg)
	}

}