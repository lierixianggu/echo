package server

import (
	"broker/api"
	"broker/utils"
	"fmt"
	"strconv"
)

//消息处理模块的实现

type MsgHandle struct {
	//存放每个MsgID 所对应的处理方法
	Apis map[uint32]api.IRouter
	//负责Worker取任务的消息队列
	TaskQueue []chan api.IRequest
	//业务工作Worker池的worker数量
	WorkerPoolSize uint32
}

//NewMsgHandle 初始化/创建MsgHandle方法
func NewMsgHandle() *MsgHandle {
	return &MsgHandle{
		Apis:           make(map[uint32]api.IRouter),
		WorkerPoolSize: utils.GlobalObject.WorkerPoolSize,
		TaskQueue:      make([]chan api.IRequest, utils.GlobalObject.MaxWorkerTaskLen),
	}
}

//DoMsgHandler 调度/执行对应的Router消息处理方法
func (mh *MsgHandle) DoMsgHandler(request api.IRequest) {
	//1.从request中找到msgID
	//  通过msgID得到router
	handler, ok := mh.Apis[request.GetMsgID()]
	if !ok {
		fmt.Println("api msgID = ", request.GetMsgID(), " is NOT FOUND! Need Register!")
	}

	//2.根据MsgID，调度对应的router业务
	handler.Handle(request)
}

//AddRouter 为消息添加具体的处理逻辑
func (mh *MsgHandle) AddRouter(msgID uint32, router api.IRouter) {
	//1.判断当前msg绑定的API是否存在
	if _, ok := mh.Apis[msgID]; ok {
		//id已经注册了
		panic("repeat api, msgID = " + strconv.Itoa(int(msgID)))
	}

	//2.添加msg与API的关系
	mh.Apis[msgID] = router
	fmt.Println("Add api MsgID = ", msgID, " succ!")
}

//StartWorkPool 启动一个Worker工作池(开启工作池的动作只能发生一次)
func (mh *MsgHandle) StartWorkPool() {
	//根据WorkerPoolSize 分别开启Worker，每个Worker用一个go来承载
	for i := 0; i < int(mh.WorkerPoolSize); i++ {
		//一个worker被启动
		//1. 当前的worker对应的channel消息队列 开辟空间 第i个worker就用第i个channel
		mh.TaskQueue[i] = make(chan api.IRequest, utils.GlobalObject.MaxWorkerTaskLen)
		//2.启动当前的Worker，阻塞等待消息从channel传递进来
		go mh.StartOneWorker(i, mh.TaskQueue[i])
	}
}

//StartOneWorker 启动一个Worker工作流程
func (mh *MsgHandle) StartOneWorker(workerID int, taskQueue chan api.IRequest) {
	fmt.Println("worker ID = ", workerID, "  is starting...")

	//不断地阻塞等待对应消息队列的消息
	for {
		select {
		//如果有消息过来，出列的就是一个客户端的Request，执行当前Request所绑定的业务
		case request := <-taskQueue:
			mh.DoMsgHandler(request)
		}
	}
}

//SendMsgToTaskQueue 将消息交给TaskQueue，由worker进行处理
func (mh *MsgHandle) SendMsgToTaskQueue(request api.IRequest) {
	//1.将消息平均分配给不同的worker
	//根据客户端建立的ConnID来进行分配
	workerID := request.GetConnection().GetConnID() % mh.WorkerPoolSize
	fmt.Println("Add ConnID = ", request.GetConnection().GetConnID(),
		" request MsgID = ", request.GetMsgID(), " to WorkerID = ", workerID)

	//2.将消息发送给对应的worker的TaskQueue
	mh.TaskQueue[workerID] <- request
}
