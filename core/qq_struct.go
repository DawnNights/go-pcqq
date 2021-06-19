package core

type QQ_Struct struct {
    StrQQ string
    BinQQ []byte
    LongQQ int64
    Utf8QQ []byte

    Time []byte
    NickName string

    RandHead16 []byte
    SessionKey []byte
    ClientKey []byte
    TgtKey []byte
    ShareKey []byte
    PublicKey []byte
    // privateKey []byte
    // sKey string

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
    // mRequstID int
}