package utils

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"io/ioutil"
	"math/rand"
	"regexp"
	"strings"
	"time"
)

//取指定长度的随机字节集
func GetRandomBin(length int) []byte {
	var bin []byte
	rand.Seed(time.Now().Unix())
	for i := 0; i < length; i++ {
		bin = append(bin,byte(rand.Intn(256)))
	}
	return bin
}

//合并字节集
func BytesMerge(args ...[]byte) []byte {
	data := bytes.Join(args,[]byte{})
	return data
}

// int16转字节集
func Int16ToBytes(num int16) []byte {
	var buffer bytes.Buffer
	binary.Write(&buffer, binary.BigEndian, num)
	return buffer.Bytes()
}

// 字节集转int16
func BytesToInt16(src []byte) int16 {
	var num int16
	bin_buf := bytes.NewBuffer(src)
	binary.Read(bin_buf, binary.BigEndian, &num)
	return num
}

// 字节集到十六进制文本
func Bin2HexTo(src []byte) string {
	s := strings.ToUpper(hex.EncodeToString(src))
	reg := regexp.MustCompile(".{2}")
	return strings.Join(reg.FindAllString(s,-1)," ")
}

// 十六进制文本到字节集
func Hex2Bin(s string) []byte {
	dst,_ := hex.DecodeString(strings.Replace(s," ","",-1))
	return dst
}

// int64到字节集
func Int64ToBytes(i int64) []byte {
    var buf = make([]byte, 8)
    binary.BigEndian.PutUint64(buf, uint64(i))
    return buf
}

// 字节集到int64
func BytesToInt64(buf []byte) int64 {
    return int64(binary.BigEndian.Uint64(buf))
}

// 反转字节集
func Flip(buf []byte) []byte {
	for i, j := 0, len(buf)-1; i < j; i, j = i+1, j-1 {
		buf[i], buf[j] = buf[j], buf[i]
	}
	return buf
}

//写入文件数据
func FileWrite(path string,content []byte) int {
	ioutil.WriteFile(path,content,0644)
	return len(content)
}