package core

import (
	"fmt"
	"pcqq/utils"
	"time"
)

type PCQQ struct {
	QQ     utils.QQ_Struct
	Client NetClient
}

// 初始化参数，连接服务器
func (self *PCQQ) Init() {
	self.QQ.PublicKey = utils.Hex2Bin("03 94 3D CB E9 12 38 61 EC F7 AD BD E3 36 91 91 07 01 50 BE 50 39 1C D3 32")
	self.QQ.ShareKey = utils.Hex2Bin("FD 0B 79 78 31 E6 88 54 FC FA EA 84 52 9C 7D 0B")
	self.QQ.RandHead16 = utils.GetRandomBin(16)
	self.Client.Connect("123.151.77.237",443)
}

// 获取登录二维码
func (self *PCQQ) GetQrCode() {
	self.Client.Send(self.pack_0825(1))
	self.unpack_0825(self.Client.Receive())

	var stateId int
	var codeId string
	var codeImg []byte

	self.Client.Send(self.pack_0818())
	self.unpack_0818(self.Client.Receive(),&codeId,&codeImg)	// 解析二维码
	utils.FileWrite("QrCode.jpg",codeImg)
	fmt.Println("ID:",codeId,"的二维码已保存至本地\n")


	for i := 0; i < 60; i++ {	// 监听扫码状态
		self.Client.Send(self.pack_0819(codeId,false))
		self.unpack_0819(self.Client.Receive(),&stateId)
		if stateId == 0 {
			self.Client.Send(self.pack_0825(1))
			self.unpack_0825(self.Client.Receive())

			self.Client.Send(self.pack_0836())
			self.unpack_0836(self.Client.Receive())

			self.Client.Send(self.pack_0828())
			self.unpack_0828(self.Client.Receive())

			self.Client.Send(self.pack_00EC(1))	// 置登录状态为上线

			self.Client.Send(self.pack_001D())
			self.unpack_001D(self.Client.Receive())

			fmt.Println("NickName:",self.QQ.NickName)
			fmt.Println("UserQQ:",self.QQ.LongQQ)
			fmt.Println("************欢迎登录************")
			return
		}
		time.Sleep(time.Second)
	}
}


// 监听消息
func (self *PCQQ) ListenMsg() {
	for {
		var data []byte = self.Client.Receive()
		if data[5] == 0 && data[6] == 23{
			self.Client.Send(self.pack_0017(self.unpack_0017(data),data[7:9]))
		}
	}
}

// 发送群消息
func (self *PCQQ) SendGroupMsg(groupId int64, content string) {
	self.Client.Send(self.pack_0002(groupId,content))
	if len(self.Client.Receive()) == 0{
		fmt.Println("Warn:","群消息发送失败")
	}else {
		fmt.Println(self.QQ.NickName+":",content)
	}
}