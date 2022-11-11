package main

import (
	"broker/server"
	"fmt"
	"io"
	"net"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:7777")

	if err != nil {
		fmt.Println("client start err, exit!")
		return
	}

	for {
		//发送封包的message消息
		dp := server.NewDataPack()
		binaryMsg, err := dp.Pack(server.NewMsgPackage(0, []byte("broker0.5 client Test Message")))

		if err != nil {
			fmt.Println("Pack error: ", err)
			return
		}

		if _, err = conn.Write(binaryMsg); err != nil {
			fmt.Println("client write error: ", err)
			return
		}

		//服务器回复了message数据

		//1. 先读取流中的head部分 得到msgID和msgDataLen
		binaryHead := make([]byte, dp.GetHeadLen())

		if _, err := io.ReadFull(conn, binaryHead); err != nil {
			fmt.Println("read head error ", err)
			break
		}
		fmt.Println(binaryHead)
		//将二进制的head拆包到msg对象中,此时只拆了head部分
		msgHead, err := dp.Unpack(binaryHead)
		if err != nil {
			fmt.Println("client unpack error: ", err)
			break
		}

		if msgHead.GetDataLen() > 0 {
			//2. 再根据DataLen进行第二次读取，将Data读出来
			//msg是一个Message对象
			msg := msgHead.(*server.Message)
			msg.Data = make([]byte, msgHead.GetDataLen())

			if _, err := io.ReadFull(conn, msg.Data); err != nil {
				fmt.Println("read msg data error: ", err)
				return
			}
			fmt.Println("recv server msg: id = ", msg.Id, " ,len = ", msg.DataLen, " ,data = ", string(msg.Data))
		}

		time.Sleep(1 * time.Second)
	}
}
