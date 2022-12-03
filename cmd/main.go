package main

import (
	"fmt"

	"github.com/clh021/event"
)

// 声明角色的结构体
type Actor struct {
}

// 为角色添加一个事件处理函数
func (a *Actor) OnEvent(param interface{}, callParam interface{}) {
	fmt.Println("actor event:", param, callParam)
}

// 全局事件
func GlobalEvent(param interface{}, callParam interface{}) {
	fmt.Println("global event:", param, callParam)
}

// 全局事件
func funcEvent(param interface{}, callParam interface{}) {
	fmt.Println("===>funcEvent:", param, callParam)
}

// 支持启动时显示构建日期和构建版本
// 需要通过命令 ` go build -ldflags "-X main.build=`git rev-parse HEAD`" ` 打包
var build = "not set"

func main() {
	fmt.Printf("Build: %s\n", build)

	e := event.New()
	// 实例化一个角色
	a := new(Actor)

	// 注册名为 `app.service1.event1` 的回调
	e.Register("app.service1.event1", event.Event{
		Priority: 300,
		Callback: a.OnEvent,
		Param:    "app.service1.event1",
	})

	// 再次在 `app.service1.event1` 上注册全局事件
	e.Register("app.service1.event2", event.Event{
		Priority: 100,
		Callback: GlobalEvent,
		Param:    22.111,
	})

	e.Register("before.app.service1.event2", event.Event{
		Priority: 100,
		Callback: funcEvent,
		Param:    "before",
	})

	e.Register("after.app.service1.event2", event.Event{
		Priority: 100,
		Callback: funcEvent,
		Param:    "after",
	})

	e.Register("app.service1.*", event.Event{
		Priority: 100,
		Callback: GlobalEvent,
		Param:    12.12,
	})

	fmt.Println("call event1:")
	// 调用事件，所有注册的同名函数都会被调用
	e.Call("app.service1.event1", 100)

	fmt.Println("call event2:")
	e.Call("app.service1.event2", 100)
}
