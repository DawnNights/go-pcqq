package core

import (
	"fmt"
	"net"
)

type NetClient struct {
	conn net.Conn
}

// 连接
func (udp *NetClient) Connect(ip string,port int16) bool {
	var err error
	udp.conn, err = net.Dial("tcp",fmt.Sprintf("%s:%d",ip,port))

	if err != nil {
        fmt.Println("Connect error:", err)
        return false
    }
    fmt.Println(fmt.Sprintf("%s:%d",ip,port),"连接成功")

    // defer udp.conn.Close()
    return true
}

// 发送
func (udp *NetClient) Send(sendData []byte) bool {
	 _, err := udp.conn.Write(sendData)
	 if err != nil {
        fmt.Println("Send error:", err)
        return false
    }
    // fmt.Println("Send Data:", utils.Bin2HexTo(sendData)+"\n")
    return true
}

// 接收
func (udp *NetClient) Receive() []byte {
	result := make([]byte, 20480)

	size,err := udp.conn.Read(result)
	if err != nil {
        fmt.Println("Receive error:", err)
        return []byte{}
    }
    // fmt.Println("Receive Data:", utils.Bin2HexTo(result[0:size])+"\n")
	return result[0:size]

}