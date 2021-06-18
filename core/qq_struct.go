package core

type QQ_Struct struct {
    binQQ []byte
    longQQ int64
    utf8QQ []byte
    int32QQ int

    pwd string
    md5 []byte
    pwdKey []byte

    time []byte
    nickName string

    randHead16 []byte
    sessionKey []byte
    clientKey []byte
    tgtKey []byte
    shareKey []byte
    publicKey []byte
    privateKey []byte
    sKey string

    pcKeyFor0819 []byte
    pcKeyTgt []byte
    pcKeyFor0828Send []byte
    pcKeyFor0828Rev []byte
    pcToken0038From0825 []byte
    pcToken0038From0818 []byte
    pcToken0060From0819 []byte
    pcToken0038From0836 []byte
    pcToken0088From0836 []byte

    localPcIp []byte
    connectSeverIp []byte
    mRequstID int
}