// 解包
package core

import (
	"fmt"
	"github.com/sun8911879/qqtea"
	"pcqq/utils"
)

// 解包_包0818: 解析二维码地址
func (self *PCQQ) decode_0818(src []byte, codeId *string,codeImg *[]byte) {
	pack :=  utils.PackDecrypt{}
	pack.SetData(src)
	pack.GetBin(16)
	dst := pack.GetAll()
	dst = dst[0:len(dst)-1]

	t,_ := qqtea.NewCipher(self.qq.ShareKey)
	dst = t.Decrypt(dst)


	pack.SetData(dst)
	pack.GetBin(7)
	self.qq.PcKeyFor0819 = pack.GetBin(16)
	pack.GetBin(4)
	length := int(pack.GetShort())

	self.qq.PcToken0038From0818 = pack.GetBin(length)
	pack.GetBin(4)
	length = int(pack.GetShort())
	*codeId = string(pack.GetBin(length))
	pack.GetBin(4)
	length = int(pack.GetShort())
	*codeImg = pack.GetBin(length)
}

/*解包_包0819: 二维码状态
stateId: 0 = 授权成功   1 = 扫码成功    2 = 未扫码   3 = 空数据包*/
func (self *PCQQ) decode_0819(src []byte, stateId *int) {
	tea,_ := qqtea.NewCipher(self.qq.PcKeyFor0819)
	pack := utils.PackDecrypt{}
	pack.SetData(src)
	pack.GetBin(16)
	dst := pack.GetAll()
	dst = dst[0:len(dst)-1]
	dst = tea.Decrypt(dst)

	pack.SetData(dst)
	*stateId = int(pack.GetByte())
	fmt.Println("二维码状态:",map[int]string{0 : "授权成功", 1 : "扫码成功", 2 : "未扫码", 3 : "空数据包"}[*stateId])
	if *stateId == 1{
		self.qq.BinQQ = src[9:13]
		self.qq.LongQQ = utils.BytesToInt64(utils.BytesMerge([]byte{0,0,0,0},self.qq.BinQQ))
		fmt.Println("当前扫码账号:",self.qq.LongQQ,"\n")
	}
	fmt.Print("\n")
	if *stateId==0 && len(dst) > 1 {
		pack.GetShort()
		length := pack.GetShort()
		self.qq.PcToken0060From0819 = pack.GetBin(int(length))
		pack.GetShort()
		length = pack.GetShort()
		self.qq.PcKeyTgt = pack.GetBin(int(length))
		// fmt.Println(utils.Bin2HexTo(self.qq.PcToken0060From0819), utils.Bin2HexTo(self.qq.PcKeyTgt))
	}
}

// 解包_包0825
func (self *PCQQ) decode_0825(src []byte) bool {
	pack := utils.PackDecrypt{}
	tea,_ := qqtea.NewCipher(self.qq.RandHead16)

	pack.SetData(src)
	pack.GetBin(16)
	dst := pack.GetAll()
	dst = dst[0:len(dst)-1]
	dst = tea.Decrypt(dst)

	pack.SetData(dst)
	Type := int(pack.GetByte())
	pack.GetShort()
	length := int(pack.GetShort())
	self.qq.PcToken0038From0825 = pack.GetBin(length)
	pack.GetBin(6)
	pack.GetBin(4)
	self.qq.LocalPcIp = pack.GetBin(4)
	pack.GetShort()

	if Type == 254{	//需要重定向服务器
		pack.GetBin(18)
		self.qq.ConnectSeverIp = pack.GetBin(4)
		fmt.Println(fmt.Sprintf("重定向地址: %d.%d.%d.%d\n",self.qq.ConnectSeverIp[0],self.qq.ConnectSeverIp[1],self.qq.ConnectSeverIp[2],self.qq.ConnectSeverIp[3]))
		return true
	}else {
		pack.GetBin(6)
		self.qq.ConnectSeverIp = pack.GetBin(4)
		return false
	}
}

// 解包_包0828
func (self *PCQQ) decode_0828(src []byte) {
	pack := utils.PackDecrypt{}
	tea,_ := qqtea.NewCipher(self.qq.PcKeyFor0828Rev)

	pack.SetData(src)
	pack.GetBin(16)
	dst := pack.GetAll()
	dst = dst[0:len(dst)-1]
	dst = tea.Decrypt(dst)

	pack.SetData(dst)
	pack.GetBin(63)
	self.qq.SessionKey = pack.GetBin(16)
}

// 解包_包0836
func (self *PCQQ) decode_0836(src []byte) bool {
	pack := utils.PackDecrypt{}
	tea,_ := qqtea.NewCipher(self.qq.ShareKey)
	tea2,_ := qqtea.NewCipher(self.qq.PcKeyTgt)

	pack.SetData(src)
	pack.GetBin(16)
	dst := pack.GetAll()
	dst = dst[0:len(dst)-1]
	if len(dst) >= 100{
		dst = tea2.Decrypt(tea.Decrypt(dst))
	}
	if len(dst)==0{return false}

	pack.SetData(dst)
	Type := int(pack.GetByte())
	if Type != 0{
		fmt.Println("08 36 返回数据TYPE出错",Type)
		return false
	}

	for i := 0; i < 5; i++ {
		tlvName := utils.Bin2HexTo(pack.GetBin(2))
		tlvLength := int(pack.GetShort())
		switch tlvName {
			case "01 09":
				pack.GetShort()
				self.qq.PcKeyFor0828Send = pack.GetBin(16)
				length := int(pack.GetShort())
				self.qq.PcToken0038From0836 = pack.GetBin(length)
				length = int(pack.GetShort())
				pack.GetBin(length)
				pack.GetShort()

			case "01 07":
				pack.GetBin(26)
				self.qq.PcKeyFor0828Rev = pack.GetBin(16)
				length := int(pack.GetShort())
				self.qq.PcToken0088From0836 = pack.GetBin(length)
				length = tlvLength - 180
				pack.GetBin(length)

			case "01 08":
				pack.GetBin(8)
				length := int(pack.GetByte())
				self.qq.NickName = string(pack.GetBin(length))
			default:
				pack.GetBin(tlvLength)
			}
	}
	if len(self.qq.PcToken0088From0836) != 136 || len(self.qq.PcKeyFor0828Send) != 16 || len(self.qq.PcToken0038From0836) != 56{
		return false
	}else {
		return true
	}
}

// 解包_包0017
func (self *PCQQ) decode_0017(src []byte) []byte {
	pack := utils.PackDecrypt{}
	tea,_ := qqtea.NewCipher(self.qq.SessionKey)
	pack.SetData(src)
	pack.GetBin(16)
	dst := pack.GetAll()
	dst = dst[0:len(dst)-1]
	dst = tea.Decrypt(dst)
	if len(dst) == 0{
		return []byte{}
	}
	info := Info{}
	pack.SetData(dst)

	info.fromGroup = pack.GetLong()	//接收群号
	info.selfQQ = pack.GetLong()	//selfQQ
	pack.GetBin(10)
	info.typeOf = pack.GetShort()	//消息类型

	length := utils.Int64ToInt(pack.GetLong())
	pack.GetBin(length)

	if len(pack.GetAll()) < 5{
		return []byte{}
	}

	pack.GetLong()
	info.flag = pack.GetByte()

	if info.typeOf == 82 && info.flag == 1{
		info.fromQQ = pack.GetLong()	//接收QQ
		info.msg.msgIndex = pack.GetLong()	//消息索引
		info.receiveTime = pack.GetLong()	//接收时间
		pack.GetBin(24)
		info.sendTime = pack.GetLong()	//发送时间
		info.msg.msgId = pack.GetLong()	//消息ID
		pack.GetBin(8)

		length = int(pack.GetShort())
		info.font = pack.GetBin(length)	//字体
		pack.GetBin(2)

		info.msg.Parse(pack.GetAll())
		fmt.Println(fmt.Sprintf("收到(%d)消息 %s[%d]: %s",info.fromGroup,info.msg.fromUserName,info.fromQQ,info.msg.msgStr))

	}
	return dst[0:16]
}

// 解包_包001D: 解析Clientkey
func (self *PCQQ) decode_001D(src []byte){
	pack := utils.PackDecrypt{}
	tea,_ := qqtea.NewCipher(self.qq.SessionKey)

	pack.SetData(src)
	pack.GetBin(9)

	pack.GetLong()  // QQ账号
	pack.GetBin(3)  // 00 00 00


	data := pack.GetAll()
	data = data[0:len(data)-1]
	data = tea.Decrypt(data)

	pack.SetData(data)
	pack.GetBin(2)


	self.qq.ClientKey = pack.GetAll()

	if len(self.qq.ClientKey) != 32{
		self.qq.ClientKey = []byte{}
		return
	}

	fmt.Println("Sessionkey",utils.Bin2HexTo(self.qq.SessionKey))
	fmt.Println("Clientkey",utils.Bin2HexTo(self.qq.ClientKey))
}