package core

import (
	"encoding/json"
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
	self.qq.PublicKey = []byte{ 3, 148, 61, 203, 233, 18, 56, 97, 236, 247, 173, 189, 227, 54, 145, 145, 7, 1, 80, 190, 80, 57, 28, 211, 50 }
	self.qq.ShareKey = []byte{ 253, 11, 121, 120, 49, 230, 136, 84, 252, 250, 234, 132, 82, 156, 125, 11 }
	self.qq.RandHead16 = []byte{ 255, 117, 107, 118, 18, 133, 105, 165, 63, 198, 146, 171, 232, 58, 175, 103 }
	self.qq.Time = utils.Int64ToBytes(time.Now().Unix())[4:]

	self.client.Connect("123.151.77.237",443)
}

// 获取二维码
func (self *PCQQ) GetQrCode() {
	data := self.touch_Send(self.encode_0825(false))
	self.decode_0825(data)

	data = self.touch_Send(self.encode_0818())
	self.checkQrCode(data)
}

// 加载本地登录信息
func (self *PCQQ)  LoadConfig() {
	config := utils.FileRead("config.json")
	json.Unmarshal(config,&self.qq)
}

// 保存当前登录信息
func (self *PCQQ) SaveConfig() {
	config,_ := json.Marshal(self.qq)
	utils.FileWrite("config.json",config)
}

// 监听QQ消息
func (self *PCQQ) ListenMessage() {
	for {
		var data []byte
		// table := map[string]bool{}

		data = self.client.Receive()
		// fmt.Println(utils.Bin2HexTo(data))
		if data[5] == 0 && data[6] == 23{
			// if _,ok := table[utils.Bin2HexTo(data)];ok {continue}else {table[utils.Bin2HexTo(data)]=true}
			self.touch_Send(self.encode_0017(self.decode_0017(data),data[7:9]))

		}
	}
}

// 发送群消息
func (self *PCQQ) SendGroupMsg(groupId int64, content string) {
	data := self.touch_Send(self.encode_0002_SendGroupText(groupId,content))
	if len(data) == 0{
		fmt.Println("<发送失败>")
	}else {
		fmt.Println(content)
	}
}

// 检查二维码状态
func (self *PCQQ) checkQrCode(src []byte) {
	var stateId int
	var codeId string
	var codeImg []byte

	self.decode_0818(src,&codeId,&codeImg)
	mod.FileWrite("QrCode.jpg",codeImg)
	fmt.Println("ID:",codeId,"的二维码已保存至本地\n")


	for i := 0; i < 60; i++ {
		src = self.touch_Send(self.encode_0819(codeId,false))
		self.decode_0819(src,&stateId)
		if stateId == 0{
			self.startLogin()
			return
		}
		time.Sleep(time.Second)
	}
	fmt.Println("您已超时，请重新执行程序")
}

// 开始登录
func (self *PCQQ) startLogin() {
	data := self.touch_Send(self.encode_0825(true))
	// fmt.Println("0825登录包发送完成")
	self.decode_0825(data)
	// fmt.Println("0825登录包解析完成")


	data = self.touch_Send(self.encode_0836())
	// fmt.Println("0836登录包发送完成")

	if !self.decode_0836(data){
		fmt.Println("0836包解析失败")
		return
	}else {
		data = self.touch_Send(self.encode_0828())
		self.decode_0828(data)

		data = self.touch_Send(self.encode_00EC(1))
		if len(data) == 0{
			fmt.Println("00EC包解析失败")
			return
		}else {
			data = self.touch_Send(self.encode_001D())
			self.decode_001D(data)
			self.qq.StrQQ = fmt.Sprintf("%d",self.qq.LongQQ)
			self.qq.Utf8QQ = []byte(self.qq.StrQQ)

		}

		fmt.Println("NickName:",self.qq.NickName)
		fmt.Println("UserQQ:",self.qq.LongQQ)
		fmt.Println("************欢迎登录************")
		self.SaveConfig()
	}

}

// 通讯_发包
func (self *PCQQ) touch_Send (sendData []byte) []byte {
	length := int16(len(sendData) + 2)
	sendData = utils.BytesMerge(utils.Int16ToBytes(length),sendData)
	self.client.Send(sendData)
	return self.client.Receive()
}


