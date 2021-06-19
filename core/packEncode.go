// 组包
package core

import (
	"github.com/sun8911879/qqtea"
	"pcqq/utils"
	"time"
)

// 组包_包0818: 获取二维码
func (self *PCQQ) encode_0818() []byte {
	self.qq.RandHead16 = utils.GetRandomBin(16)
	tea,_ := qqtea.NewCipher(self.qq.RandHead16)
	data := utils.BytesMerge(
		[]byte{2,54,57,8,24,179,166,0,0,0,0,3,0,0,0,1,1,1,0,0,103,183,0,0,0,0},
		self.qq.RandHead16,
		tea.Encrypt([]byte{0,25,0,16,0,1,0,0,4,76,0,0,0,1,0,0,21,81,0,0,1,20,0,29,1,2,0,25,3,148,61,203,233,18,56,97,236,247,173,189,227,54,145,145,7,1,80,190,80,57,28,211,50,3,5,0,30,0,0,0,0,0,0,0,5,0,0,0,4,0,0,0,0,0,0,0,72,0,0,0,2,0,0,0,2,0,0}),
		[]byte{3},
		)
	return data
}

// 组包_包0819: 二维码状态
// login: true = 授权登录    false = 取二维码验证状态
func (self *PCQQ) encode_0819(codeId string, login bool) []byte {
	var data []byte
	self.qq.RandHead16 = utils.GetRandomBin(16)
	tea,_ := qqtea.NewCipher(self.qq.PcKeyFor0819)

	pack := utils.PackEncrypt{}
	pack.SetHex("00 19 00 10 00 01 00 00 04 4C 00 00 00 01 00 00 15 51 00 00 03 01 00 22")
	pack.SetShort(int16(len(codeId)))
	pack.SetStr(codeId)
	if login{
		pack.SetHex("03 14 00 02 00 00")
	}
	data = tea.Encrypt(pack.GetAll())

	pack.Empty()
	pack.SetHex("02 36 39")
	pack.SetHex("08 19")
	pack.SetBin(utils.GetRandomBin(2))
	pack.SetHex("00 00 00 00 03 00 00 00 01 01 01 00 00 67 B7 00 00 00 00 00 30 00 3A")
	pack.SetShort(int16(len(self.qq.PcToken0038From0818)))
	pack.SetBin(self.qq.PcToken0038From0818)
	pack.SetBin(data)
	pack.SetHex("03")
	data = pack.GetAll()

	return data
}

// 组包_包0825
// login: true = 二维码登录	false = 取二维码
func (self *PCQQ) encode_0825(login bool) []byte {
	var data []byte
	self.qq.RandHead16 = utils.GetRandomBin(16)
	tea,_ := qqtea.NewCipher(self.qq.RandHead16)
	if login{
		pack := utils.PackEncrypt{}

		pack.SetHex("00 18 00 16 00 01 00 00 04 4C 00 00 00 01 00 00 15 51")
		pack.SetBin(self.qq.BinQQ)
		pack.SetHex("00 00 00 00 03 09 00 08 00 01")
		pack.SetBin(self.qq.ConnectSeverIp)
		pack.SetHex("00 01 00 36 00 12 00 02 00 01 00 00 00 00 00 00 00 00 00 00 00 00 00 00 01 14 00 1D")
		pack.SetHex("01 02")
		pack.SetShort(int16(len(self.qq.PublicKey)))
		pack.SetBin(self.qq.PublicKey)
		data = tea.Encrypt(pack.GetAll())

		pack.Empty()
		pack.SetHex("02 36 39")
		pack.SetHex("08 25")
		pack.SetBin(utils.GetRandomBin(2))
		pack.SetBin(self.qq.BinQQ)
		pack.SetHex("03 00 00 00 01 01 01 00 00 67 B7 00 00 00 00")
		pack.SetBin(self.qq.RandHead16)
		pack.SetBin(data)
		pack.SetHex("03")
		data = pack.GetAll()
	}else {
		data = utils.BytesMerge(
		[]byte{2,54,57,8,37,41,35,0,0,0,0,3,0,0,0,1,1,1,0,0,103,183,0,0,0,0},
		self.qq.RandHead16,
		tea.Encrypt([]byte{ 0, 24, 0, 22, 0, 1, 0, 0, 4, 76, 0, 0, 0, 1, 0, 0, 21, 81, 0, 0, 0, 0, 0, 0, 0, 0, 0, 4, 0, 12, 0, 0, 0, 8, 113, 114, 95, 108, 111, 103, 105, 110, 3, 9, 0, 8, 0, 1, 0, 0, 0, 0, 0, 4, 1, 20, 0, 29, 1, 2, 0, 25, 3, 148, 61, 203, 233, 18, 56, 97, 236, 247, 173, 189, 227, 54, 145, 145, 7, 1, 80, 190, 80, 57, 28, 211, 50 }),
		[]byte{3},
		)
	}

	return data
}

// 组包_包0828
func (self *PCQQ) encode_0828() []byte {
	tea,_ := qqtea.NewCipher(self.qq.PcKeyFor0828Send)
    tlv := utils.Tlv{}
    pack := utils.PackEncrypt{}

    pack.Empty()
    pack.SetBin(tlv.Tlv007(self.qq.PcToken0088From0836))
    pack.SetBin(tlv.Tlv00C(self.qq.ConnectSeverIp))
    pack.SetBin(tlv.Tlv015())
    pack.SetBin(tlv.Tlv036())
    pack.SetBin(tlv.Tlv018(self.qq.BinQQ))
    pack.SetBin(tlv.Tlv01F())
    pack.SetBin(tlv.Tlv105())
    pack.SetBin(tlv.Tlv10B())
    pack.SetBin(tlv.Tlv02D())
    data := tea.Encrypt(pack.GetAll())

    pack.Empty()
    pack.SetHex("02 36 39")
    pack.SetHex("08 28")
    pack.SetBin(utils.GetRandomBin(2))
    pack.SetBin(self.qq.BinQQ)
    pack.SetHex("02 00 00 00 01 01 01 00 00 67 B7 00 30 00 3A")
    pack.SetShort(int16(len((self.qq.PcToken0038From0836))))
    pack.SetBin(self.qq.PcToken0038From0836)
    pack.SetBin(data)
    pack.SetHex("03")
    data = pack.GetAll()
    return data
}

// 组包_包0836
func (self *PCQQ) encode_0836() []byte {
	tlv := utils.Tlv{}
	tea,_ := qqtea.NewCipher(self.qq.ShareKey)
	pack := utils.PackEncrypt{}

	self.qq.RandHead16 = utils.GetRandomBin(16)
	pack.Empty()
	pack.SetBin(tlv.Tlv112(self.qq.PcToken0038From0825))
	pack.SetBin(tlv.Tlv30F(pcName))
	pack.SetBin(tlv.Tlv005(self.qq.BinQQ))
	pack.SetBin(tlv.Tlv303(self.qq.PcToken0060From0819))
	pack.SetBin(tlv.Tlv015())
	pack.SetBin(tlv.Tlv01A(self.qq.PcKeyTgt))
	pack.SetBin(tlv.Tlv018(self.qq.BinQQ))
	pack.SetBin(tlv.Tlv103())
	pack.SetBin(tlv.Tlv312())
	pack.SetBin(tlv.Tlv313())
	pack.SetBin(tlv.Tlv102(self.qq.PcToken0038From0825))
	data := tea.Encrypt(pack.GetAll())

	pack.Empty()
	pack.SetHex("02 36 39")
	pack.SetHex("08 36")
	pack.SetBin(utils.GetRandomBin(2))
	pack.SetBin(self.qq.BinQQ)
	pack.SetHex("03 00 00 00 01 01 01 00 00 67 B7 00 00 00 00 00 01")
	pack.SetHex("01 02")
	pack.SetShort(int16(len(self.qq.PublicKey)))
	pack.SetBin(self.qq.PublicKey)
	pack.SetHex("00 00 00 10")
	pack.SetBin(self.qq.RandHead16)
	pack.SetBin(data)
	pack.SetHex("03")
	data = pack.GetAll()
	return data
}

// 组包_包00EC
func (self *PCQQ) encode_00EC(state int) []byte{
	tea,_ := qqtea.NewCipher(self.qq.SessionKey)
	pack := utils.PackEncrypt{}
	pack.SetHex("01 00")
	switch state {
	case 1:
		pack.SetHex("0A")
	case 2:
		pack.SetHex("3C")
	case 3:
		pack.SetHex("1E")
	case 4:
		pack.SetHex("32")
	case 5:
		pack.SetHex("46")
	case 6:
		pack.SetHex("28")
	default:
		pack.SetHex("0A")
	}
	pack.SetHex("00 01 00 01 00 04 00 00 00 00")
	data := tea.Encrypt(pack.GetAll())
	pack.Empty()
	pack.SetHex("02 36 39")
	pack.SetHex("00 EC")
	pack.SetBin(utils.GetRandomBin(2))
	pack.SetBin(self.qq.BinQQ)
	pack.SetHex("02 00 00 00 01 01 01 00 00 67 B7")
	pack.SetBin(data)
	pack.SetHex("03")
	data = pack.GetAll()
	return data
}

// 组包_包0017: 申请已读
func (self *PCQQ) encode_0017(sendData []byte, sequence []byte) []byte{
	pack := utils.PackEncrypt{}
	tea,_ := qqtea.NewCipher(self.qq.SessionKey)

	pack.Empty()
	pack.SetBin(sendData)
	data := tea.Encrypt(pack.GetAll())

	pack.Empty()
	pack.SetHex("02 36 39")
	pack.SetHex("00 17")
	pack.SetBin(sequence)
	pack.SetBin(self.qq.BinQQ)
	pack.SetHex("02 00 00 00 01 01 01 00 00 67 B7")
	pack.SetBin(data)

	pack.SetHex("03")
	data = pack.GetAll()
	return data
}

// 组包_包001D: 更新Clientkey
func (self *PCQQ) encode_001D() []byte {
	tea,_ := qqtea.NewCipher(self.qq.SessionKey)
	pack := utils.PackEncrypt{}

	pack.Empty()
	pack.SetHex("11")
	data := tea.Encrypt(pack.GetAll())

	pack.Empty()
	pack.SetHex("02 36 39")
	pack.SetHex("00 1D")
	pack.SetBin(utils.GetRandomBin(2))
	pack.SetBin(self.qq.BinQQ)
	pack.SetHex("02 00 00 00 01 01 01 00 00 67 B7")
	pack.SetBin(data)
	pack.SetHex("03")
	data = pack.GetAll()
	return data
}

// 组包_包0058: 心跳包
func (self *PCQQ) encode_0058() []byte {
	tea,_ := qqtea.NewCipher(self.qq.SessionKey)
	pack := utils.PackEncrypt{}

	pack.Empty()
	pack.SetBin(self.qq.Utf8QQ)
	data := tea.Encrypt(pack.GetAll())

	pack.Empty()
	pack.SetHex("02 36 39")
	pack.SetHex("00 58")
	pack.SetBin(utils.GetRandomBin(2))
	pack.SetBin(self.qq.BinQQ)
	pack.SetHex("02 00 00 00 01 01 01 00 00 67 B7")
	pack.SetBin(data)
	pack.SetHex("03")
	data = pack.GetAll()
	return data
}

// 组包_包0002QQ群文本消息
func (self *PCQQ) encode_0002_SendGroupText(groupId int64, content string) []byte {
	tea,_ := qqtea.NewCipher(self.qq.SessionKey)
	pack := utils.PackEncrypt{}

	self.qq.Time = utils.Int64ToBytes(time.Now().Unix())[4:]
	Msg := []byte(content)
	// Msg = Msg[0:len(Msg)-1]

	pack.Empty()
	pack.SetHex("00 01 01 00 00 00 00 00 00 00 4D 53 47 00 00 00 00 00")
	pack.SetBin(self.qq.Time)
	pack.SetBin(utils.Flip(self.qq.Time))
	pack.SetHex("00 00 00 00 09 00 86 00 00 06 E5 AE 8B E4 BD 93 00 00 01")
	pack.SetShort(int16(len(Msg) + 3))
	pack.SetHex("01")
	pack.SetShort(int16(len(Msg)))
	pack.SetBin(Msg)
	data := pack.GetAll()

	pack.Empty()
	pack.SetHex("2A")
	pack.SetLong(groupId)
	pack.SetShort(int16(len(data)))
	pack.SetBin(data)
	data = tea.Encrypt(pack.GetAll())


	pack.Empty()
	pack.SetHex("02 36 39")
	pack.SetHex("00 02")
	pack.SetBin(utils.GetRandomBin(2))
	pack.SetBin(self.qq.BinQQ)
	pack.SetHex("02 00 00 00 01 01 01 00 00 67 B7")
	pack.SetBin(data)
	pack.SetHex("03")
	data = pack.GetAll()
	// fmt.Println("消息包构建完成",len(data))
	return data
}