package api

import "net"

//IConnection 定义连接模块的抽象层
type IConnection interface {
	//Start 启动连接 让当前连接开始工作
	Start()
	//Stop 停止连接 结束当前连接的工作
	Stop()
	//GetTCPConnection 获取当前连接的绑定socket conn
	GetTCPConnection() *net.TCPConn
	//GetConnID 获取当前连接模块的连接ID
	GetConnID() uint32
	//RemoteAddr 获取对端的TCP状态 IP port
	RemoteAddr() net.Addr
	//SendMsg 发送数据，将数据发送给远程的客户端
	SendMsg(msgId uint32, data []byte) error
}

//HandleFunc 定义一个处理连接业务的方法
type HandleFunc func(*net.TCPConn, []byte, int) error
