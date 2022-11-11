package server

type Message struct {
	//消息的id
	Id uint32
	//消息的长度
	DataLen uint32
	//消息的内容
	Data []byte
}

//NewMsgPackage 创建一个Message消息包
func NewMsgPackage(id uint32, data []byte) *Message {
	return &Message{
		Id:      id,
		DataLen: uint32(len(data)),
		Data:    data,
	}
}

//GetMsgId 获取消息的ID
func (msg *Message) GetMsgId() uint32 {
	return msg.Id
}

//GetDataLen 获取消息的长度
func (msg *Message) GetDataLen() uint32 {
	return msg.DataLen
}

//GetData 获取消息的内容
func (msg *Message) GetData() []byte {
	return msg.Data
}

//SetMsgId 设计消息的ID
func (msg *Message) SetMsgId(id uint32) {
	msg.Id = id
}

//SetData 设计消息内容
func (msg *Message) SetData(data []byte) {
	msg.Data = data
}

//SetDataLen 设置数据段长度
func (msg *Message) SetDataLen(len uint32) {
	msg.DataLen = len
}
