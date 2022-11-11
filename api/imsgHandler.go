package api

type IMsgHandle interface {
	//DoMsgHandler 调度/执行对应的Router消息处理方法
	DoMsgHandler(request IRequest)
	//AddRouter 为消息添加具体的处理逻辑
	AddRouter(msgID uint32, router IRouter)
	//StartWorkPool 启动一个Worker工作池(开启工作池的动作只能发生一次)
	StartWorkPool()
	//StartOneWorker 启动一个Worker工作流程
	StartOneWorker(workerID int, taskQueue chan IRequest)
	//SendMsgToTaskQueue 将消息交给TaskQueue，由worker进行处理
	SendMsgToTaskQueue(request IRequest)
}
