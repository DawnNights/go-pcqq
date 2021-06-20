package core

import (
	"fmt"
	"net"
	"pcqq/utils"
)

type NetClient struct {
	conn net.Conn
}

// 连接
func (self *NetClient) Connect(ip string,port int16) bool {
	var err error
	self.conn, err = net.Dial("tcp",fmt.Sprintf("%s:%d",ip,port))
	if err != nil {
        fmt.Println("Connect error:", err)
        return false
    }
    fmt.Println(fmt.Sprintf("%s:%d",ip,port),"连接成功")
    return true
}

// 发送
func (self *NetClient) Send(sendData []byte) bool {
	length := int16(len(sendData) + 2)
	sendData = utils.BytesMerge(utils.Int16ToBytes(length),sendData)

	 _, err := self.conn.Write(sendData)
	 if err != nil {
        fmt.Println("Send error:", err)
        return false
    }
    // fmt.Println("Send Data:", utils.Bin2HexTo(sendData)+"\n")
    return true
}

// 接收
func (self *NetClient) Receive() []byte {
	result := make([]byte, 20480)

	size,err := self.conn.Read(result)
	if err != nil {
        fmt.Println("Receive error:", err)
        return []byte{}
    }
    // fmt.Println("Receive Data:", utils.Bin2HexTo(result[0:size])+"\n")
	return result[0:size]

}