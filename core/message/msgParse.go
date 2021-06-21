// 消息解析
package message

import (
	"pcqq/utils"
	"strings"
)

func MessageParse(body []byte,msg *string) {
	unpack := utils.PackDecrypt{}
	unpack.SetData(body)

	msgType := unpack.GetByte()	//消息类型
	dataLen := int(unpack.GetShort())	//数据长度
	pos := unpack.GetPosition()

	for pos+dataLen < len(unpack.GetAll()) {
		unpack.GetByte()
		switch msgType {
		case 1:	// 纯文本 && 艾特
			length := int(unpack.GetShort())
			str := string(unpack.GetBin(length))
			if strings.Index(str,"@") == 0 && pos+dataLen-unpack.GetPosition() == 16 {
				unpack.GetBin(10)
				*msg = *msg + str
				unpack.GetLong()
				
			}else {
				*msg = *msg + str
			}
		case 2:	// Emoji(系统表情)
			unpack.GetShort()
			*msg = *msg + FaceToStr(int(unpack.GetByte()))
		case 3:	// 图片
			length := int(unpack.GetShort())
			picStr := string(unpack.GetBin(length))
			*msg = *msg + "[Image:" + picStr + "]"
		}

		unpack.GetBin(pos+dataLen-unpack.GetPosition())
		msgType = unpack.GetByte()
		dataLen = int(unpack.GetShort())
		pos = unpack.GetPosition()
	}

}