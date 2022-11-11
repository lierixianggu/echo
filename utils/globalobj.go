package utils

import (
	"broker/api"
	"encoding/json"
	"io/ioutil"
)

//存储全局参数

type GlobalObj struct {
	//server对象
	TcpServer api.IServer
	//服务器主机监听的ip
	Host string
	//当前服务器port
	TcpPort int
	//当前服务器名称
	Name string
	//当前服务器版本
	Version string
	//当前服务器允许最大连接数
	MaxConn int
	//当前服务器数据包的最大值
	MaxPackageSize uint32
	//当前业务工作Worker池的Goroutine数量
	WorkerPoolSize uint32
	//每个worker对应的消息队列的任务的数量最大值
	MaxWorkerTaskLen uint32
}

//GlobalObject 定义一个对外的全局对象
var GlobalObject *GlobalObj

func (g *GlobalObj) Reload() {
	data, err := ioutil.ReadFile("../../conf/broker.json")
	if err != nil {
		panic(err)
	}

	//将json数据解析到Object中
	err = json.Unmarshal(data, &GlobalObject)
	if err != nil {
		panic(err)
	}
}

//init 用来初始化当前的GlobalObject
func init() {
	//配置默认值
	GlobalObject = &GlobalObj{
		Name:             "ServerApp",
		Version:          "broker0.1",
		TcpPort:          7777,
		Host:             "0.0.0.0",
		MaxConn:          1000,
		MaxPackageSize:   4096,
		WorkerPoolSize:   10,
		MaxWorkerTaskLen: 1024,
	}
	GlobalObject.Reload()
}
