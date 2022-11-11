package api

type IDataPack interface {
	//GetHeadLen 获取头长度
	GetHeadLen() uint32
	//Pack 装包
	Pack(msg IMessage) ([]byte, error)
	//Unpack 拆包
	Unpack([]byte) (IMessage, error)
}
