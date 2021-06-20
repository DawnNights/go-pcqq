// 组装协议包
package core

import (
	"pcqq/utils"
	"time"
)

// packType=1 取二维码	packType=2 二维码登录
func (self *PCQQ) pack_0825(packType int) []byte {
	var data []byte
	self.QQ.RandHead16 = utils.GetRandomBin(16)
	pack := utils.PackEncrypt{}
	tea,_ := utils.NewCipher(self.QQ.RandHead16)

	switch packType {
	case 1:
		pack.SetHex("00 18 00 16 00 01 00 00 04 4C 00 00 00 01 00 00 15 51 00 00 00 00 00 00 00 00 00 04 00 0C 00 00 00 08 71 72 5F 6C 6F 67 69 6E 03 09 00 08 00 01 00 00 00 00 00 04 01 14 00 1D")
		pack.SetHex("01 02")
		pack.SetShort(int16(len((self.QQ.PublicKey))))
		pack.SetBin(self.QQ.PublicKey)
		data = tea.Encrypt(pack.GetAll())

		pack.Empty()
		pack.SetHex("02 36 39")
		pack.SetHex("08 25")
		pack.SetBin(utils.GetRandomBin(2))
		pack.SetHex("00 00 00 00 03 00 00 00 01 01 01 00 00 67 B7 00 00 00 00")
		pack.SetBin(self.QQ.RandHead16)
		pack.SetBin(data)
		pack.SetHex("03")
		data = pack.GetAll()
	case 2:
		pack.SetHex("00 18 00 16 00 01 00 00 04 4C 00 00 00 01 00 00 15 51")
		pack.SetBin(self.QQ.BinQQ)
		pack.SetHex("00 00 00 00 03 09 00 08 00 01")
		pack.SetBin(self.QQ.ConnectSeverIp)
		pack.SetHex("00 01 00 36 00 12 00 02 00 01 00 00 00 00 00 00 00 00 00 00 00 00 00 00 01 14 00 1D")
		pack.SetHex("01 02")
		pack.SetShort(int16(len(self.QQ.PublicKey)))
		pack.SetBin(self.QQ.PublicKey)
		data = tea.Encrypt(pack.GetAll())

		pack.Empty()
		pack.SetHex("02 36 39")
		pack.SetHex("08 25")
		pack.SetBin(utils.GetRandomBin(2))
		pack.SetBin(self.QQ.BinQQ)
		pack.SetHex("03 00 00 00 01 01 01 00 00 67 B7 00 00 00 00")
		pack.SetBin(self.QQ.RandHead16)
		pack.SetBin(data)
		pack.SetHex("03")
		data = pack.GetAll()
	}

	return data
}

// 获取二维码
func (self *PCQQ) pack_0818() []byte {
	self.QQ.RandHead16 = utils.GetRandomBin(16)
	pack := utils.PackEncrypt{}
	tea,_ := utils.NewCipher(self.QQ.RandHead16)

	pack.Empty()
	pack.SetHex("00 19 00 10 00 01 00 00 04 4C 00 00 00 01 00 00 15 51 00 00 01 14 00 1D")
	pack.SetHex("01 02")
	pack.SetShort(int16(len(self.QQ.PublicKey)))
	pack.SetBin(self.QQ.PublicKey)
	pack.SetHex("03 05 00 1E 00 00 00 00 00 00 00 05 00 00 00 04 00 00 00 00 00 00 00 48 00 00 00 02 00 00 00 02 00 00")
	data := tea.Encrypt(pack.GetAll())

	pack.Empty()
	pack.SetHex("02 36 39")
	pack.SetHex("08 18")
	pack.SetBin(utils.GetRandomBin(2))
	pack.SetHex("00 00 00 00 03 00 00 00 01 01 01 00 00 67 B7 00 00 00 00")
	pack.SetBin(self.QQ.RandHead16)
	pack.SetBin(data)
	pack.SetHex("03")
	data = pack.GetAll()
	return data
}

// 检查二维码状态
// codeId 二维码ID
// login: true = 授权登录    false = 取二维码验证状态
func (self *PCQQ) pack_0819(codeId string, login bool) []byte {
	var data []byte
	self.QQ.RandHead16 = utils.GetRandomBin(16)
	tea,_ := utils.NewCipher(self.QQ.PcKeyFor0819)

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
	pack.SetShort(int16(len(self.QQ.PcToken0038From0818)))
	pack.SetBin(self.QQ.PcToken0038From0818)
	pack.SetBin(data)
	pack.SetHex("03")
	data = pack.GetAll()

	return data
}

func (self *PCQQ) pack_0836() []byte {
	tlv := utils.Tlv{}
	tea,_ := utils.NewCipher(self.QQ.ShareKey)
	pack := utils.PackEncrypt{}

	self.QQ.RandHead16 = utils.GetRandomBin(16)
	pack.Empty()
	pack.SetBin(tlv.Tlv112(self.QQ.PcToken0038From0825))
	pack.SetBin(tlv.Tlv30F("DawnNights"))
	pack.SetBin(tlv.Tlv005(self.QQ.BinQQ))
	pack.SetBin(tlv.Tlv303(self.QQ.PcToken0060From0819))
	pack.SetBin(tlv.Tlv015())
	pack.SetBin(tlv.Tlv01A(self.QQ.PcKeyTgt))
	pack.SetBin(tlv.Tlv018(self.QQ.BinQQ))
	pack.SetBin(tlv.Tlv103())
	pack.SetBin(tlv.Tlv312())
	pack.SetBin(tlv.Tlv313())
	pack.SetBin(tlv.Tlv102(self.QQ.PcToken0038From0825))
	data := tea.Encrypt(pack.GetAll())

	pack.Empty()
	pack.SetHex("02 36 39")
	pack.SetHex("08 36")
	pack.SetBin(utils.GetRandomBin(2))
	pack.SetBin(self.QQ.BinQQ)
	pack.SetHex("03 00 00 00 01 01 01 00 00 67 B7 00 00 00 00 00 01")
	pack.SetHex("01 02")
	pack.SetShort(int16(len(self.QQ.PublicKey)))
	pack.SetBin(self.QQ.PublicKey)
	pack.SetHex("00 00 00 10")
	pack.SetBin(self.QQ.RandHead16)
	pack.SetBin(data)
	pack.SetHex("03")
	data = pack.GetAll()
	return data
}

func (self *PCQQ) pack_0828() []byte {
	tea,_ := utils.NewCipher(self.QQ.PcKeyFor0828Send)
    tlv := utils.Tlv{}
    pack := utils.PackEncrypt{}

    pack.Empty()
    pack.SetBin(tlv.Tlv007(self.QQ.PcToken0088From0836))
    pack.SetBin(tlv.Tlv00C(self.QQ.ConnectSeverIp))
    pack.SetBin(tlv.Tlv015())
    pack.SetBin(tlv.Tlv036())
    pack.SetBin(tlv.Tlv018(self.QQ.BinQQ))
    pack.SetBin(tlv.Tlv01F())
    pack.SetBin(tlv.Tlv105())
    pack.SetBin(tlv.Tlv10B())
    pack.SetBin(tlv.Tlv02D())
    data := tea.Encrypt(pack.GetAll())

    pack.Empty()
    pack.SetHex("02 36 39")
    pack.SetHex("08 28")
    pack.SetBin(utils.GetRandomBin(2))
    pack.SetBin(self.QQ.BinQQ)
    pack.SetHex("02 00 00 00 01 01 01 00 00 67 B7 00 30 00 3A")
    pack.SetShort(int16(len((self.QQ.PcToken0038From0836))))
    pack.SetBin(self.QQ.PcToken0038From0836)
    pack.SetBin(data)
    pack.SetHex("03")
    data = pack.GetAll()
    return data
}

// 更新Clientkey
func (self *PCQQ) pack_001D() []byte {
	tea,_ := utils.NewCipher(self.QQ.SessionKey)
	pack := utils.PackEncrypt{}

	pack.Empty()
	pack.SetHex("11")
	data := tea.Encrypt(pack.GetAll())

	pack.Empty()
	pack.SetHex("02 36 39")
	pack.SetHex("00 1D")
	pack.SetBin(utils.GetRandomBin(2))
	pack.SetBin(self.QQ.BinQQ)
	pack.SetHex("02 00 00 00 01 01 01 00 00 67 B7")
	pack.SetBin(data)
	pack.SetHex("03")
	data = pack.GetAll()
	return data
}

// state: 上线状态    1 = 在线   2 = Q我吧   3 = 离开   4 = 忙碌   5 = 请勿打扰   6 = 隐身
func (self *PCQQ) pack_00EC(state int) []byte{
	tea,_ := utils.NewCipher(self.QQ.SessionKey)
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
	pack.SetBin(self.QQ.BinQQ)
	pack.SetHex("02 00 00 00 01 01 01 00 00 67 B7")
	pack.SetBin(data)
	pack.SetHex("03")
	data = pack.GetAll()
	return data
}

// 确认消息已读
// sendData=消息包解密前16位
// sequence=消息包序列
func (self *PCQQ) pack_0017(sendData []byte, sequence []byte) []byte{
	pack := utils.PackEncrypt{}
	tea,_ := utils.NewCipher(self.QQ.SessionKey)

	pack.Empty()
	pack.SetBin(sendData)
	data := tea.Encrypt(pack.GetAll())

	pack.Empty()
	pack.SetHex("02 36 39")
	pack.SetHex("00 17")
	pack.SetBin(sequence)
	pack.SetBin(self.QQ.BinQQ)
	pack.SetHex("02 00 00 00 01 01 01 00 00 67 B7")
	pack.SetBin(data)

	pack.SetHex("03")
	data = pack.GetAll()
	return data
}

// 发送群文本消息包
// groupId=发送群号
// content=消息内容
func (self *PCQQ) pack_0002(groupId int64, content string) []byte {
	tea,_ := utils.NewCipher(self.QQ.SessionKey)
	pack := utils.PackEncrypt{}

	self.QQ.Time = utils.Int64ToBytes(time.Now().Unix())[4:]
	Msg := []byte(content)

	pack.Empty()
	pack.SetHex("00 01 01 00 00 00 00 00 00 00 4D 53 47 00 00 00 00 00")
	pack.SetBin(self.QQ.Time)
	pack.SetBin(utils.Flip(self.QQ.Time))
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
	pack.SetBin(self.QQ.BinQQ)
	pack.SetHex("02 00 00 00 01 01 01 00 00 67 B7")
	pack.SetBin(data)
	pack.SetHex("03")
	data = pack.GetAll()
	return data
}

 /*
// 心跳包
func (self *PCQQ) pack_0058() []byte {
	tea,_ := utils.NewCipher(self.QQ.SessionKey)
	pack := utils.PackEncrypt{}

	pack.Empty()
	pack.SetBin(self.QQ.Utf8QQ)
	data := tea.Encrypt(pack.GetAll())

	pack.Empty()
	pack.SetHex("02 36 39")
	pack.SetHex("00 58")
	pack.SetBin(utils.GetRandomBin(2))
	pack.SetBin(self.QQ.BinQQ)
	pack.SetHex("02 00 00 00 01 01 01 00 00 67 B7")
	pack.SetBin(data)
	pack.SetHex("03")
	data = pack.GetAll()
	return data
}
 */