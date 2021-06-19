package utils

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"io/ioutil"
	"math/rand"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unsafe"
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

//取出中间文本
func StrMidGet(pre string,suf string,str string) string {
	n := strings.Index(str, pre)
	if n == -1 {n = 0} else {n = n + len(pre)}
	str = string([]byte(str)[n:])
	m := strings.Index(str, suf)
	if m == -1 {m = len(str)}
	return string([]byte(str)[:m])
}

//写入文件数据
func FileWrite(path string,content []byte) int {
	ioutil.WriteFile(path,content,0644)
	return len(content)
}

//读取文件数据
func FileRead(path string) []byte {
	res,_ := os.Open(path)
	defer res.Close()
	data,_ := ioutil.ReadAll(res)
	return data
}

//int64转int
func Int64ToInt(num int64) int {
	intPtr := (*int)(unsafe.Pointer(&num))
	v := *intPtr
	return v
}

// string转int
func StrToInt(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}