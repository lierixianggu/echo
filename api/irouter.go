package api

//路由的抽象接口，路由里的数据都是IRequest

type IRouter interface {
	//Handle 处理conn业务的钩子方法
	Handle(request IRequest)
}
