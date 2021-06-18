package core

import (
	"fmt"
	"mod"
	"pcqq/utils"
	"time"
)

const pcName = "DawnNights"
type PCQQ struct {
	client NetClient
	qq     QQ_Struct
}

// 初始化
func (self *PCQQ) Init() {
	self.qq.publicKey = []byte{ 3, 148, 61, 203, 233, 18, 56, 97, 236, 247, 173, 189, 227, 54, 145, 145, 7, 1, 80, 190, 80, 57, 28, 211, 50 }
	self.qq.shareKey = []byte{ 253, 11, 121, 120, 49, 230, 136, 84, 252, 250, 234, 132, 82, 156, 125, 11 }
	self.qq.randHead16 = []byte{ 255, 117, 107, 118, 18, 133, 105, 165, 63, 198, 146, 171, 232, 58, 175, 103 }
	self.qq.time = utils.Int64ToBytes(time.Now().Unix())[4:]

	self.client.Connect("123.151.77.237",443)
}

// 获取二维码
func (self *PCQQ) GetQrCode() {
	data := self.Touch_Send(self.Encode_0825(false))
	self.Decode_0825(data)

	data = self.Touch_Send(self.Encode_0818())
	self.checkQrCode(data)
}

// 检查二维码状态
func (self *PCQQ) checkQrCode(src []byte) {
	var stateId int
	var codeId string
	var codeImg []byte

	self.Decode_0818(src,&codeId,&codeImg)
	mod.FileWrite("QrCode.jpg",codeImg)
	fmt.Println("ID:",codeId,"的二维码已保存至本地\n")


	for i := 0; i < 60; i++ {
		src = self.Touch_Send(self.Encode_0819(codeId,false))
		self.Decode_0819(src,&stateId)
		if stateId == 0{
			self.Login()
			return
		}
		time.Sleep(time.Second)
	}
	fmt.Println("您已超时，请重新执行程序")
}

// 开始登录
func (self *PCQQ) Login() {
	data := self.Touch_Send(self.Encode_0825(true))
	// fmt.Println("0825登录包发送完成")
	self.Decode_0825(data)
	// fmt.Println("0825登录包解析完成")


	data = self.Touch_Send(self.Encode_0836())
	// fmt.Println("0836登录包发送完成")

	if !self.Decode_0836(data){
		fmt.Println("0836包解析失败")
		return
	}else {
		data = self.Touch_Send(self.Encode_0828())
		self.Decode_0828(data)

		data = self.Touch_Send(self.Encode_00EC(1))
		if len(data) == 0{
			fmt.Println("00EC包解析失败")
			return
		}else {
			data = self.Touch_Send(self.Encode_001D())
			self.Decode_001D(data)
			self.qq.utf8QQ = []byte(fmt.Sprintf("%d",self.qq.longQQ))
		}

		fmt.Println("NickName:",self.qq.nickName)
		fmt.Println("UserQQ:",self.qq.longQQ)
		fmt.Println("************欢迎登录************")

	}

}

// 通讯_发包
func (self *PCQQ) Touch_Send (sendData []byte) []byte {
	length := int16(len(sendData) + 2)
	sendData = utils.BytesMerge(utils.Int16ToBytes(length),sendData)
	self.client.Send(sendData)
	return self.client.Receive()
}

// 发送群消息
func (self *PCQQ) SendGroupMsg(groupId int64, content string) {
	data := self.Touch_Send(self.Encode_0002_SendGroupText(groupId,content))
	if len(data) == 0{
		fmt.Println("<发送失败>")
	}else {
		fmt.Println(content)
	}
}



