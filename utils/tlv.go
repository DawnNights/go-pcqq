package utils

type Tlv [4]uint32

func(t *Tlv) TlvPack(TlvCmd string,TlvBin []byte) []byte {
	var pack PackEncrypt
	var data []byte
	pack.Empty()
	pack.SetHex(TlvCmd)
	pack.SetShort(int16(len(TlvBin)))
	pack.SetBin(TlvBin)
	data = pack.GetAll()
	return data
}



func(t *Tlv) Tlv112(Token0038From0825 []byte) []byte{
	var pack PackEncrypt
	pack.Empty()
	pack.SetBin(Token0038From0825)
	return t.TlvPack("01 12",pack.GetAll())
}


func(t *Tlv) Tlv30F (PcName string) []byte{
	var pack PackEncrypt
	pack.Empty()
	pack.SetShort(int16(len(PcName)))
	pack.SetStr(PcName)
	return t.TlvPack("03 0F",pack.GetAll())
}


func(t *Tlv) Tlv005(BinQQ []byte) []byte{
	var pack PackEncrypt
	pack.Empty()
	pack.SetHex("00 02")
	pack.SetBin(BinQQ)
	return t.TlvPack("00 05",pack.GetAll())
}


func(t *Tlv) Tlv303(PcToken0060From0819 []byte) []byte{
	var pack PackEncrypt
	pack.Empty()
	pack.SetBin(PcToken0060From0819)
	return t.TlvPack("03 03",pack.GetAll())
}


func(t *Tlv) Tlv015() []byte{
	var pack PackEncrypt
	pack.Empty()
	pack.SetHex("00 01 01 74 83 F2 C3 00 10 14 FE 77 FC 00 00 00 00 00 00 00 00 00 00 00 00 02 17 65 6E 9D 00 10 78 8A 33 DD 00 76 A1 78 EB 8E 5B BB FF 17 D0 10")
	return t.TlvPack("00 15",pack.GetAll())
}


func(t *Tlv) Tlv01A(Tgtkey []byte) []byte{
	var pack PackEncrypt
	tea,_ := NewCipher(Tgtkey)
	pack.Empty()
	pack.SetHex("00 01 01 74 83 F2 C3 00 10 14 FE 77 FC 00 00 00 00 00 00 00 00 00 00 00 00 02 17 65 6E 9D 00 10 78 8A 33 DD 00 76 A1 78 EB 8E 5B BB FF 17 D0 10")
	return t.TlvPack("00 1A",tea.Encrypt(pack.GetAll()))
}


func(t *Tlv) Tlv018(BinQQ []byte) []byte{
	var pack PackEncrypt
	pack.Empty()
	pack.SetHex("00 01 00 00 04 4C 00 00 00 01 00 00 15 51")
	pack.SetBin(BinQQ)
	pack.SetHex("00 00 00 00")
	return t.TlvPack("00 18",pack.GetAll())
}


func(t *Tlv) Tlv103() []byte{
	var pack PackEncrypt
	pack.Empty()
	pack.SetHex("00 01 00 10")
	pack.SetBin(GetRandomBin(16))
	return t.TlvPack("01 03",pack.GetAll())
}


func(t *Tlv) Tlv312() []byte{
	var pack PackEncrypt
	pack.Empty()
	pack.SetHex("01 00 00 00 00")
	return(t.TlvPack("03 12",pack.GetAll()))
}


func(t *Tlv) Tlv313() []byte{
	var pack PackEncrypt
	pack.Empty()
	pack.SetHex("01 01 02 00 10 EE 47 7F A4 BC D6 EE 65 02 65 4D E9 43 38 4C 3D 00 00 00 EB")
	return t.TlvPack("03 13",pack.GetAll())
}


func(t *Tlv) Tlv102(Token0038From0825 []byte) []byte{
	var pack PackEncrypt
	pack.Empty()
	pack.SetHex("00 01")
	pack.SetBin(GetRandomBin(16))
	pack.SetShort(int16(len(Token0038From0825)))
	pack.SetBin(Token0038From0825)
	pack.SetHex("00 14")
	pack.SetBin(GetRandomBin(20))
	return t.TlvPack("01 02",pack.GetAll())
}


func(t *Tlv) Tlv007(PcToken0088From0836 []byte) []byte{
	var pack PackEncrypt
	pack.Empty()
	pack.SetBin(PcToken0088From0836)
	return t.TlvPack("00 07",pack.GetAll())
}


func(t *Tlv) Tlv00C(ServerIp []byte) []byte{
	var pack PackEncrypt
	pack.Empty()
	pack.SetHex("00 02 00 01 00 00 00 00 00 00 00 00")
	pack.SetBin(ServerIp)
	pack.SetHex("00 50 00 00 00 00")
	return t.TlvPack("00 0C",pack.GetAll())
}

func(t *Tlv) Tlv036() []byte{
	var pack PackEncrypt
	pack.Empty()
	pack.SetHex("00 02 00 01 00 00 00 00 00 00 00 00 00 00 00 00 00 00")
	return t.TlvPack("00 36",pack.GetAll())
}

func(t *Tlv) Tlv01F() []byte{
	var pack PackEncrypt
	pack.Empty()
	pack.SetHex("00 01")
	pack.SetBin(GetRandomBin(32))
	return t.TlvPack("00 1F",pack.GetAll())
}


func(t *Tlv) Tlv105() []byte{
	var pack PackEncrypt
	pack.Empty()
	pack.SetHex("00 01 01 02 00 14 01 01 00 10")
	pack.SetBin(GetRandomBin(16))
	pack.SetHex("00 14 01 02 00 10")
	pack.SetBin(GetRandomBin(16))
	return t.TlvPack("01 05",pack.GetAll())
}


func(t *Tlv) Tlv10B() []byte{
	var pack PackEncrypt
	pack.Empty()
	pack.SetHex("00 02")
	pack.SetBin(GetRandomBin(17))
	pack.SetHex("10 00 00 00 00 00 00 00 02 00 63 3E 00 63 02 04 00 03 07 00 04 00 49 F5 00 00 00 00 78 8A 33 DD 00 76 A1 78 EB 8E 5B BB FF 17 D0 10 01 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 01 00 00 00 01 00 00 00 01 00 00 00 01 00 FE 26 81 75 EC 2A 34 EF 02 3E 50 39 6D B1 AF CC 9F EA 54 E1 70 CC 6C 9E 4E 63 8B 51 EC 7C 84 5C 68 00 00 00 00")
	return t.TlvPack("01 0B",pack.GetAll())
}


func(t *Tlv) Tlv02D() []byte{
	var pack PackEncrypt
	pack.Empty()
	pack.SetHex("00 01")
	pack.SetHex("C0 A8 74 83")  // 192.168.116.131
	return t.TlvPack("00 2D",pack.GetAll())
}
