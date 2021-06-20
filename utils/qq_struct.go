package utils

type QQ_Struct struct {
    BinQQ []byte
    LongQQ int64

    Time []byte
    NickName string

    RandHead16 []byte
    SessionKey []byte
    ClientKey []byte
    ShareKey []byte
    PublicKey []byte

    PcKeyFor0819 []byte
    PcKeyTgt []byte
    PcKeyFor0828Send []byte
    PcKeyFor0828Rev []byte
    PcToken0038From0825 []byte
    PcToken0038From0818 []byte
    PcToken0060From0819 []byte
    PcToken0038From0836 []byte
    PcToken0088From0836 []byte

    LocalPcIp []byte
    ConnectSeverIp []byte
}