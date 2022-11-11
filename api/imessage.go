package api

//将请求的消息封装到Message中

type IMessage interface {
	//GetMsgId 获取消息的ID
	GetMsgId() uint32
	//GetDataLen 获取消息的长度
	GetDataLen() uint32
	//GetData 获取消息的内容
	GetData() []byte
	//SetMsgId 设计消息的ID
	SetMsgId(uint32)
	//SetData 设计消息内容
	SetData([]byte)
	//SetDataLen 设置数据段长度
	SetDataLen(uint32)
}
