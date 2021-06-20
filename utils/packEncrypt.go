package utils

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