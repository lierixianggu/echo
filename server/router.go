package server

import "broker/api"

type BaseRouter struct {
}

//Handle 处理conn业务的钩子方法
func (br *BaseRouter) Handle(request api.IRequest) {

}
