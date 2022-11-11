package main

import (
	"broker/api"
	"broker/server"
	"fmt"
)

type PingRouter struct {
	server.BaseRouter
}

func (this *PingRouter) Handle(request api.IRequest) {
	fmt.Println("Call ping Router Handle")
	//先读取客户端的数据，再写回ping..
	fmt.Println("recv from client:msgID = ", request.GetMsgID(), " , data = ", string(request.GetData()))
	err := request.GetConnection().SendMsg(1, []byte("ping...ping...ping..."))
	if err != nil {
		fmt.Println(err)
	}
}

type HelloRouter struct {
	server.BaseRouter
}

func (this *HelloRouter) Handle(request api.IRequest) {
	fmt.Println("Call hello Router Handle")
	//先读取客户端的数据，再写回ping..
	fmt.Println("recv from client:msgID = ", request.GetMsgID(), " , data = ", string(request.GetData()))
	err := request.GetConnection().SendMsg(0, []byte("hello...hello...hello..."))
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	//1.创建一个server句柄
	s := server.NewServer("[broker 0.6]")
	//2.给当前服务器添加一个自定义的router
	s.AddRouter(1, &PingRouter{})
	s.AddRouter(0, &HelloRouter{})
	//2.启动server
	s.Serve()
}
