// 解密协议包
package core

import (
	"fmt"
	"pcqq/utils"
	"strings"
	"time"
)

func (self *PCQQ) unpack_0825(src []byte) bool {
	unpack := utils.PackDecrypt{}
	tea,_ := utils.NewCipher(self.QQ.RandHead16)

	unpack.SetData(src)
	unpack.GetBin(16)
	dst := unpack.GetAll()
	dst = dst[0:len(dst)-1]
	dst = tea.Decrypt(dst)

	unpack.SetData(dst)
	Type := int(unpack.GetByte())
	unpack.GetShort()
	length := int(unpack.GetShort())
	self.QQ.PcToken0038From0825 = unpack.GetBin(length)
	unpack.GetBin(6)
	unpack.GetBin(4)
	self.QQ.LocalPcIp = unpack.GetBin(4)
	unpack.GetShort()

	if Type == 254{	//需要重定向服务器
		unpack.GetBin(18)
		self.QQ.ConnectSeverIp = unpack.GetBin(4)
		fmt.Println(fmt.Sprintf(
			"重定向地址: %d.%d.%d.%d\n",
			self.QQ.ConnectSeverIp[0],
			self.QQ.ConnectSeverIp[1],
			self.QQ.ConnectSeverIp[2],
			self.QQ.ConnectSeverIp[3],
			))
		return true
	}else {
		unpack.GetBin(6)
		self.QQ.ConnectSeverIp = unpack.GetBin(4)
		return false
	}
}

// 解析二维码地址
func (self *PCQQ) unpack_0818(src []byte, codeId *string,codeImg *[]byte) {
	unpack :=  utils.PackDecrypt{}
	unpack.SetData(src)
	unpack.GetBin(16)
	dst := unpack.GetAll()
	dst = dst[0:len(dst)-1]

	t,_ := utils.NewCipher(self.QQ.ShareKey)
	dst = t.Decrypt(dst)


	unpack.SetData(dst)
	unpack.GetBin(7)
	self.QQ.PcKeyFor0819 = unpack.GetBin(16)
	unpack.GetBin(4)
	length := int(unpack.GetShort())

	self.QQ.PcToken0038From0818 = unpack.GetBin(length)
	unpack.GetBin(4)
	length = int(unpack.GetShort())
	*codeId = string(unpack.GetBin(length))
	unpack.GetBin(4)
	length = int(unpack.GetShort())
	*codeImg = unpack.GetBin(length)
}

// 二维码状态
// stateId: 0 = 授权成功   1 = 扫码成功    2 = 未扫码   3 = 空数据包
func (self *PCQQ) unpack_0819(src []byte, stateId *int) {
	tea,_ := utils.NewCipher(self.QQ.PcKeyFor0819)
	unpack := utils.PackDecrypt{}
	unpack.SetData(src)
	unpack.GetBin(16)
	dst := unpack.GetAll()
	dst = dst[0:len(dst)-1]
	dst = tea.Decrypt(dst)

	unpack.SetData(dst)
	*stateId = int(unpack.GetByte())
	fmt.Println("二维码状态:",map[int]string{0 : "授权成功", 1 : "扫码成功", 2 : "未扫码", 3 : "空数据包"}[*stateId])

	if *stateId == 1{
		self.QQ.BinQQ = src[9:13]
		self.QQ.LongQQ = utils.BytesToInt64(utils.BytesMerge([]byte{0,0,0,0},self.QQ.BinQQ))
		fmt.Println("当前扫码账号:",self.QQ.LongQQ)
	}
	fmt.Print("\n")

	if *stateId==0 && len(dst) > 1 {
		unpack.GetShort()
		length := unpack.GetShort()
		self.QQ.PcToken0060From0819 = unpack.GetBin(int(length))
		unpack.GetShort()
		length = unpack.GetShort()
		self.QQ.PcKeyTgt = unpack.GetBin(int(length))
	}
}

func (self *PCQQ) unpack_0836(src []byte) bool {
	unpack := utils.PackDecrypt{}
	tea,_ := utils.NewCipher(self.QQ.ShareKey)
	tea2,_ := utils.NewCipher(self.QQ.PcKeyTgt)

	unpack.SetData(src)
	unpack.GetBin(16)
	dst := unpack.GetAll()
	dst = dst[0:len(dst)-1]
	if len(dst) >= 100{
		dst = tea2.Decrypt(tea.Decrypt(dst))
	}
	if len(dst)==0{return false}

	unpack.SetData(dst)
	Type := int(unpack.GetByte())
	if Type != 0{
		fmt.Println("08 36 返回数据TYPE出错",Type)
		return false
	}

	for i := 0; i < 5; i++ {
		tlvName := utils.Bin2HexTo(unpack.GetBin(2))
		tlvLength := int(unpack.GetShort())
		switch tlvName {
			case "01 09":
				unpack.GetShort()
				self.QQ.PcKeyFor0828Send = unpack.GetBin(16)
				length := int(unpack.GetShort())
				self.QQ.PcToken0038From0836 = unpack.GetBin(length)
				length = int(unpack.GetShort())
				unpack.GetBin(length)
				unpack.GetShort()

			case "01 07":
				unpack.GetBin(26)
				self.QQ.PcKeyFor0828Rev = unpack.GetBin(16)
				length := int(unpack.GetShort())
				self.QQ.PcToken0088From0836 = unpack.GetBin(length)
				length = tlvLength - 180
				unpack.GetBin(length)

			case "01 08":
				unpack.GetBin(8)
				length := int(unpack.GetByte())
				self.QQ.NickName = string(unpack.GetBin(length))
			default:
				unpack.GetBin(tlvLength)
			}
	}
	if len(self.QQ.PcToken0088From0836) != 136 || len(self.QQ.PcKeyFor0828Send) != 16 || len(self.QQ.PcToken0038From0836) != 56{
		return false
	}else {
		return true
	}
}

// 解析SessionKey
func (self *PCQQ) unpack_0828(src []byte) {
	unpack := utils.PackDecrypt{}
	tea,_ := utils.NewCipher(self.QQ.PcKeyFor0828Rev)

	unpack.SetData(src)
	unpack.GetBin(16)
	dst := unpack.GetAll()
	dst = dst[0:len(dst)-1]
	dst = tea.Decrypt(dst)

	unpack.SetData(dst)
	unpack.GetBin(63)
	self.QQ.SessionKey = unpack.GetBin(16)
}

// 解析Clientkey
func (self *PCQQ) unpack_001D(src []byte){
	unpack := utils.PackDecrypt{}
	tea,_ := utils.NewCipher(self.QQ.SessionKey)

	unpack.SetData(src)
	unpack.GetBin(9)

	unpack.GetLong()  // QQ账号
	unpack.GetBin(3)  // 00 00 00


	data := unpack.GetAll()
	data = data[0:len(data)-1]
	data = tea.Decrypt(data)

	unpack.SetData(data)
	unpack.GetBin(2)


	self.QQ.ClientKey = unpack.GetAll()

	if len(self.QQ.ClientKey) != 32{
		self.QQ.ClientKey = []byte{}
		return
	}

	fmt.Println("Clientkey",utils.Bin2HexTo(self.QQ.ClientKey))
	fmt.Println("Sessionkey",utils.Bin2HexTo(self.QQ.SessionKey))
}

// 解析消息包
func (self *PCQQ) unpack_0017(src []byte) []byte {
	unpack := utils.PackDecrypt{}
	tea,_ := utils.NewCipher(self.QQ.SessionKey)
	unpack.SetData(src)
	unpack.GetBin(16)
	dst := unpack.GetAll()
	dst = dst[0:len(dst)-1]
	dst = tea.Decrypt(dst)
	if len(dst) == 0{
		return []byte{}
	}

	unpack.SetData(dst)

	fromGroup := unpack.GetLong()	//接收群号
	unpack.GetLong()	//自身QQ
	unpack.GetBin(10)
	typeOf := unpack.GetShort()	//消息类型

	length := utils.Int64ToInt(unpack.GetLong())
	unpack.GetBin(length)

	if len(unpack.GetAll()) < 5{
		return []byte{}
	}

	unpack.GetLong()
	flag := unpack.GetByte()

	if typeOf == 82 && flag == 1{
		fromQQ := unpack.GetLong()	//接收QQ
		unpack.GetLong()	//消息索引
		receiveTime := unpack.GetLong()	//接收时间
		unpack.GetBin(24)
		unpack.GetLong()	//发送时间
		unpack.GetLong()	//消息ID
		unpack.GetBin(8)

		length = int(unpack.GetShort())
		unpack.GetBin(length)	//字体
		unpack.GetBin(2)

		msgType := unpack.GetByte()
		unpack.GetShort()	//数据长度
		unpack.GetByte()

		if msgType != 1 {
			return []byte{}
		}

		length := int(unpack.GetShort())
		msgStr := string(unpack.GetBin(length))
		unpack.GetBin(2)


		str := utils.Bin2HexTo(unpack.GetAll())
		str = str[strings.Index(str,"04 00 C0 04 00 CA 04 00 F8 04 00")+105:]
		unpack.SetData(utils.Hex2Bin(str))

		length = int(unpack.GetShort())
		fromUserName := string(unpack.GetBin(length))

		fmt.Println(fmt.Sprintf(
			"<%s>收到(%d)消息 %s[%d]: %s",
			time.Unix(receiveTime,0).Format("2006-01-02 03:04:05"),
			fromGroup,fromUserName,fromQQ,msgStr,
			))

	}
	return dst[0:16]
}