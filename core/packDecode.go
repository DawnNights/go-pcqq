// 解包
package core

import (
	"fmt"
	"github.com/sun8911879/qqtea"
	"pcqq/utils"
)

// 解包_包0818: 解析二维码地址
func (self *PCQQ) Decode_0818(src []byte, codeId *string,codeImg *[]byte) {
	pack :=  utils.PackDecrypt{}
	pack.SetData(src)
	pack.GetBin(16)
	dst := pack.GetAll()
	dst = dst[0:len(dst)-1]

	t,_ := qqtea.NewCipher(self.qq.shareKey)
	dst = t.Decrypt(dst)


	pack.SetData(dst)
	pack.GetBin(7)
	self.qq.pcKeyFor0819 = pack.GetBin(16)
	pack.GetBin(4)
	length := int(pack.GetShort())

	self.qq.pcToken0038From0818 = pack.GetBin(length)
	pack.GetBin(4)
	length = int(pack.GetShort())
	*codeId = string(pack.GetBin(length))
	pack.GetBin(4)
	length = int(pack.GetShort())
	*codeImg = pack.GetBin(length)
}

/*解包_包0819: 二维码状态
stateId: 0 = 授权成功   1 = 扫码成功    2 = 未扫码   3 = 空数据包*/
func (self *PCQQ) Decode_0819(src []byte, stateId *int) {
	tea,_ := qqtea.NewCipher(self.qq.pcKeyFor0819)
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
		self.qq.binQQ = src[9:13]
		self.qq.longQQ = utils.BytesToInt64(utils.BytesMerge([]byte{0,0,0,0},self.qq.binQQ))
		fmt.Println("当前扫码账号:",self.qq.longQQ,"\n")
	}
	fmt.Print("\n")
	if *stateId==0 && len(dst) > 1 {
		pack.GetShort()
		length := pack.GetShort()
		self.qq.pcToken0060From0819 = pack.GetBin(int(length))
		pack.GetShort()
		length = pack.GetShort()
		self.qq.pcKeyTgt = pack.GetBin(int(length))
		// fmt.Println(utils.Bin2HexTo(self.qq.pcToken0060From0819), utils.Bin2HexTo(self.qq.pcKeyTgt))
	}
}

// 解包_包0825
func (self *PCQQ) Decode_0825(src []byte) bool {
	pack := utils.PackDecrypt{}
	tea,_ := qqtea.NewCipher(self.qq.randHead16)

	pack.SetData(src)
	pack.GetBin(16)
	dst := pack.GetAll()
	dst = dst[0:len(dst)-1]
	dst = tea.Decrypt(dst)

	pack.SetData(dst)
	Type := int(pack.GetByte())
	pack.GetShort()
	length := int(pack.GetShort())
	self.qq.pcToken0038From0825 = pack.GetBin(length)
	pack.GetBin(6)
	pack.GetBin(4)
	self.qq.localPcIp = pack.GetBin(4)
	pack.GetShort()

	if Type == 254{	//需要重定向服务器
		pack.GetBin(18)
		self.qq.connectSeverIp = pack.GetBin(4)
		return true
	}else {
		pack.GetBin(6)
		self.qq.connectSeverIp = pack.GetBin(4)
		return false
	}
}

// 解包_包0828
func (self *PCQQ) Decode_0828(src []byte) {
	pack := utils.PackDecrypt{}
	tea,_ := qqtea.NewCipher(self.qq.pcKeyFor0828Rev)

	pack.SetData(src)
	pack.GetBin(16)
	dst := pack.GetAll()
	dst = dst[0:len(dst)-1]
	dst = tea.Decrypt(dst)

	pack.SetData(dst)
	pack.GetBin(63)
	self.qq.sessionKey = pack.GetBin(16)
}

// 解包_包0836
func (self *PCQQ) Decode_0836(src []byte) bool {
	pack := utils.PackDecrypt{}
	tea,_ := qqtea.NewCipher(self.qq.shareKey)
	tea2,_ := qqtea.NewCipher(self.qq.pcKeyTgt)

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
				self.qq.pcKeyFor0828Send = pack.GetBin(16)
				length := int(pack.GetShort())
				self.qq.pcToken0038From0836 = pack.GetBin(length)
				length = int(pack.GetShort())
				pack.GetBin(length)
				pack.GetShort()

			case "01 07":
				pack.GetBin(26)
				self.qq.pcKeyFor0828Rev = pack.GetBin(16)
				length := int(pack.GetShort())
				self.qq.pcToken0088From0836 = pack.GetBin(length)
				length = tlvLength - 180
				pack.GetBin(length)

			case "01 08":
				pack.GetBin(8)
				length := int(pack.GetByte())
				self.qq.nickName = string(pack.GetBin(length))
			default:
				pack.GetBin(tlvLength)
			}
	}
	if len(self.qq.pcToken0088From0836) != 136 || len(self.qq.pcKeyFor0828Send) != 16 || len(self.qq.pcToken0038From0836) != 56{
		return false
	}else {
		return true
	}
}

// 解包_包001D: 解析Clientkey
func (self *PCQQ) Decode_001D(src []byte){
	pack := utils.PackDecrypt{}
	tea,_ := qqtea.NewCipher(self.qq.sessionKey)

	pack.SetData(src)
	pack.GetBin(9)
	pack.GetLong()  // QQ账号
	pack.GetBin(3)  // 00 00 00
	data := pack.GetAll()
	data = data[0:len(data)-1]
	data = tea.Encrypt(data)

	pack.SetData(data)
	pack.GetBin(2)
	self.qq.clientKey = pack.GetAll()

	fmt.Println("Sessionkey:",utils.Bin2HexTo(self.qq.sessionKey))
	// fmt.Println("QQSkey",self.qq.sKey)
	fmt.Println("Clientkey:",utils.Bin2HexTo(self.qq.clientKey))
}

