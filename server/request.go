package server

import "broker/api"

type Request struct {
	//已经和客户端建立好的连接
	conn api.IConnection
	//客户端请求的数据
	msg api.IMessage
}

//GetConnection 得到当前连接
func (r *Request) GetConnection() api.IConnection {
	return r.conn
}

//GetData 得到请求的消息数据
func (r *Request) GetData() []byte {
	return r.msg.GetData()
}

//GetMsgID 得到请求的消息ID
func (r *Request) GetMsgID() uint32 {
	return r.msg.GetMsgId()
}
