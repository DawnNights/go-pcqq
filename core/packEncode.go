// 组包
package core

import (
	"github.com/sun8911879/qqtea"
	"pcqq/utils"
)

// 组包_包0818: 获取二维码
func (self *PCQQ) Encode_0818() []byte {
	self.qq.randHead16 = utils.GetRandomBin(16)
	tea,_ := qqtea.NewCipher(self.qq.randHead16)
	data := utils.BytesMerge(
		[]byte{2,54,57,8,24,179,166,0,0,0,0,3,0,0,0,1,1,1,0,0,103,183,0,0,0,0},
		self.qq.randHead16,
		tea.Encrypt([]byte{0,25,0,16,0,1,0,0,4,76,0,0,0,1,0,0,21,81,0,0,1,20,0,29,1,2,0,25,3,148,61,203,233,18,56,97,236,247,173,189,227,54,145,145,7,1,80,190,80,57,28,211,50,3,5,0,30,0,0,0,0,0,0,0,5,0,0,0,4,0,0,0,0,0,0,0,72,0,0,0,2,0,0,0,2,0,0}),
		[]byte{3},
		)
	return data
}

// 组包_包0819: 二维码状态
// login: true = 授权登录    false = 取二维码验证状态
func (self *PCQQ) Encode_0819(codeId string, login bool) []byte {
	var data []byte
	self.qq.randHead16 = utils.GetRandomBin(16)
	tea,_ := qqtea.NewCipher(self.qq.pcKeyFor0819)

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
	pack.SetShort(int16(len(self.qq.pcToken0038From0818)))
	pack.SetBin(self.qq.pcToken0038From0818)
	pack.SetBin(data)
	pack.SetHex("03")
	data = pack.GetAll()

	return data
}

// 组包_包0825
// login: true = 二维码登录	false = 取二维码
func (self *PCQQ) Encode_0825(login bool) []byte {
	var data []byte
	self.qq.randHead16 = utils.GetRandomBin(16)
	tea,_ := qqtea.NewCipher(self.qq.randHead16)
	if login{
		pack := utils.PackEncrypt{}

		pack.SetHex("00 18 00 16 00 01 00 00 04 4C 00 00 00 01 00 00 15 51")
		pack.SetBin(self.qq.binQQ)
		pack.SetHex("00 00 00 00 03 09 00 08 00 01")
		pack.SetBin(self.qq.connectSeverIp)
		pack.SetHex("00 01 00 36 00 12 00 02 00 01 00 00 00 00 00 00 00 00 00 00 00 00 00 00 01 14 00 1D")
		pack.SetHex("01 02")
		pack.SetShort(int16(len(self.qq.publicKey)))
		pack.SetBin(self.qq.publicKey)
		data = tea.Encrypt(pack.GetAll())

		pack.Empty()
		pack.SetHex("02 36 39")
		pack.SetHex("08 25")
		pack.SetBin(utils.GetRandomBin(2))
		pack.SetBin(self.qq.binQQ)
		pack.SetHex("03 00 00 00 01 01 01 00 00 67 B7 00 00 00 00")
		pack.SetBin(self.qq.randHead16)
		pack.SetBin(data)
		pack.SetHex("03")
		data = pack.GetAll()
	}else {
		data = utils.BytesMerge(
		[]byte{2,54,57,8,37,41,35,0,0,0,0,3,0,0,0,1,1,1,0,0,103,183,0,0,0,0},
		self.qq.randHead16,
		tea.Encrypt([]byte{ 0, 24, 0, 22, 0, 1, 0, 0, 4, 76, 0, 0, 0, 1, 0, 0, 21, 81, 0, 0, 0, 0, 0, 0, 0, 0, 0, 4, 0, 12, 0, 0, 0, 8, 113, 114, 95, 108, 111, 103, 105, 110, 3, 9, 0, 8, 0, 1, 0, 0, 0, 0, 0, 4, 1, 20, 0, 29, 1, 2, 0, 25, 3, 148, 61, 203, 233, 18, 56, 97, 236, 247, 173, 189, 227, 54, 145, 145, 7, 1, 80, 190, 80, 57, 28, 211, 50 }),
		[]byte{3},
		)
	}

	return data
}

// 组包_包0828
func (self *PCQQ) Encode_0828() []byte {
	tea,_ := qqtea.NewCipher(self.qq.pcKeyFor0828Send)
    tlv := utils.Tlv{}
    pack := utils.PackEncrypt{}

    pack.Empty()
    pack.SetBin(tlv.Tlv007(self.qq.pcToken0088From0836))
    pack.SetBin(tlv.Tlv00C(self.qq.connectSeverIp))
    pack.SetBin(tlv.Tlv015())
    pack.SetBin(tlv.Tlv036())
    pack.SetBin(tlv.Tlv018(self.qq.binQQ))
    pack.SetBin(tlv.Tlv01F())
    pack.SetBin(tlv.Tlv105())
    pack.SetBin(tlv.Tlv10B())
    pack.SetBin(tlv.Tlv02D())
    data := tea.Encrypt(pack.GetAll())

    pack.Empty()
    pack.SetHex("02 36 39")
    pack.SetHex("08 28")
    pack.SetBin(utils.GetRandomBin(2))
    pack.SetBin(self.qq.binQQ)
    pack.SetHex("02 00 00 00 01 01 01 00 00 67 B7 00 30 00 3A")
    pack.SetShort(int16(len((self.qq.pcToken0038From0836))))
    pack.SetBin(self.qq.pcToken0038From0836)
    pack.SetBin(data)
    pack.SetHex("03")
    data = pack.GetAll()
    return data
}

// 组包_包0836
func (self *PCQQ) Encode_0836() []byte {
	tlv := utils.Tlv{}
	tea,_ := qqtea.NewCipher(self.qq.shareKey)
	pack := utils.PackEncrypt{}

	self.qq.randHead16 = utils.GetRandomBin(16)
	pack.Empty()
	pack.SetBin(tlv.Tlv112(self.qq.pcToken0038From0825))
	pack.SetBin(tlv.Tlv30F(pcName))
	pack.SetBin(tlv.Tlv005(self.qq.binQQ))
	pack.SetBin(tlv.Tlv303(self.qq.pcToken0060From0819))
	pack.SetBin(tlv.Tlv015())
	pack.SetBin(tlv.Tlv01A(self.qq.pcKeyTgt))
	pack.SetBin(tlv.Tlv018(self.qq.binQQ))
	pack.SetBin(tlv.Tlv103())
	pack.SetBin(tlv.Tlv312())
	pack.SetBin(tlv.Tlv313())
	pack.SetBin(tlv.Tlv102(self.qq.pcToken0038From0825))
	data := tea.Encrypt(pack.GetAll())

	pack.Empty()
	pack.SetHex("02 36 39")
	pack.SetHex("08 36")
	pack.SetBin(utils.GetRandomBin(2))
	pack.SetBin(self.qq.binQQ)
	pack.SetHex("03 00 00 00 01 01 01 00 00 67 B7 00 00 00 00 00 01")
	pack.SetHex("01 02")
	pack.SetShort(int16(len(self.qq.publicKey)))
	pack.SetBin(self.qq.publicKey)
	pack.SetHex("00 00 00 10")
	pack.SetBin(self.qq.randHead16)
	pack.SetBin(data)
	pack.SetHex("03")
	data = pack.GetAll()
	return data
}

// 组包_包00EC
func (self *PCQQ) Encode_00EC(state int) []byte{
	tea,_ := qqtea.NewCipher(self.qq.sessionKey)
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
	pack.SetBin(self.qq.binQQ)
	pack.SetHex("02 00 00 00 01 01 01 00 00 67 B7")
	pack.SetBin(data)
	pack.SetHex("03")
	data = pack.GetAll()
	return data
}

// 组包_包001D: 更新Clientkey
func (self *PCQQ) Encode_001D() []byte {
	tea,_ := qqtea.NewCipher(self.qq.sessionKey)
	pack := utils.PackEncrypt{}

	pack.Empty()
	pack.SetHex("11")
	data := tea.Encrypt(pack.GetAll())

	pack.Empty()
	pack.SetHex("02 36 39")
	pack.SetHex("00 1D")
	pack.SetBin(utils.GetRandomBin(2))
	pack.SetBin(self.qq.binQQ)
	pack.SetHex("02 00 00 00 01 01 01 00 00 67 B7")
	pack.SetBin(data)
	pack.SetHex("03")
	data = pack.GetAll()
	return data
}
