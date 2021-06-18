package utils

// 解包
type PackDecrypt struct {
	src []byte
}

// 置数据
func (pack *PackDecrypt) SetData(data []byte)  {
	pack.src = data
}
// 取字节集
func (pack *PackDecrypt) GetBin(length int) []byte {
	temp := pack.src[0:length]
	pack.src = pack.src[length:]
	return temp
}
// 取短整数
func (pack *PackDecrypt) GetShort() int16 {
	num := BytesToInt16(pack.src[0:2])
	pack.src = pack.src[2:]
	return num
}
// 取长整数
func (pack *PackDecrypt) GetLong() int64 {
	num := BytesToInt64(BytesMerge([]byte{0,0,0,0},pack.src[0:4]))
	pack.src = pack.src[4:]
	return num
}
// 取字节
func (pack *PackDecrypt) GetByte() byte {
	temp := pack.src[0]
	pack.src = pack.src[1:]
	return temp
}
// 取所有数据
func (pack *PackDecrypt) GetAll() []byte {
	return pack.src
}

// 打包
type PackEncrypt struct {
	src []byte
}

// 清空
func (pack *PackEncrypt) Empty() {
	pack.src = []byte{}
}
// 置字节集
func (pack *PackEncrypt) SetBin(data []byte) {
	pack.src = BytesMerge(pack.src,data)
}
// 置文本十六进制格式
func (pack *PackEncrypt) SetHex(s string) {
	pack.src = BytesMerge(pack.src,Hex2Bin(s))
}
// 置短整数
func (pack *PackEncrypt) SetShort(num int16) {
	pack.src = BytesMerge(pack.src,Int16ToBytes(num))
}
// 置长整数
func (pack *PackEncrypt) SetLong(num int64) {
	pack.src = BytesMerge(pack.src,Int64ToBytes(num)[4:])
}
// 置字符串
func (pack *PackEncrypt) SetStr(str string) {
	pack.src = BytesMerge(pack.src,[]byte(str))
}
// 取所有数据
func (pack *PackEncrypt) GetAll() []byte {
	return pack.src
}