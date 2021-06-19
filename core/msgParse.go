package core

import (
	"pcqq/utils"
	"strings"
)

type Info struct {
    msg Message // 消息体

    selfQQ int64    // 自身QQ
    fromQQ  int64   // 接收QQ
    fromGroup int64 // 接收群号

    typeOf int16
    flag byte   // 信息类型
    font []byte // 字体数据

    sendTime int64  // 发送时间戳
    receiveTime int64   // 接收时间戳

    // content []byte  //信息内容
}

type Message struct {
    msgId int64
    msgType byte
    msgIndex int64

    msgStr string
    msgLen int16

    fromUserName string
}

func (msg *Message) Parse(content []byte) {
	var pack utils.PackDecrypt
	pack.SetData(content)

	msg.msgType = pack.GetByte()
	msg.msgLen = pack.GetShort()
	pack.GetByte()

	if msg.msgType != 1 {
		return
	}

	length := int(pack.GetShort())
	msg.msgStr = string(pack.GetBin(length))
	pack.GetBin(2)

	/*pack.GetBin(7)
	pack.GetBin(11)

	length = utils.StrToInt(utils.Bin2HexTo([]byte{pack.GetByte()}))
	pack.GetBin(length)
	pack.GetBin(10)
	pack.GetBin(35)*/

	str := utils.Bin2HexTo(pack.GetAll())
	str = str[strings.Index(str,"04 00 C0 04 00 CA 04 00 F8 04 00")+105:]
	pack.SetData(utils.Hex2Bin(str))

	length = int(pack.GetShort())
	msg.fromUserName = string(pack.GetBin(length))
}