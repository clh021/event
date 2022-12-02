package main

import (
	"fmt"

	"github.com/clh021/event"
)

// 声明角色的结构体
type Actor struct {
}

// 为角色添加一个事件处理函数
func (a *Actor) OnEvent(param interface{}) {
	fmt.Println("actor event:", param)
}

// 全局事件
func GlobalEvent(param interface{}) {
	fmt.Println("global event:", param)
}

// 支持启动时显示构建日期和构建版本
// 需要通过命令 ` go build -ldflags "-X main.build=`git rev-parse HEAD`" ` 打包
var build = "not set"

func main() {
	fmt.Printf("Build: %s\n", build)

	e := event.New()
	// 实例化一个角色
	a := new(Actor)

	// 注册名为OnSkill的回调
	e.RegisterWithPriority("OnSkill", a.OnEvent, 300)

	// 再次在OnSkill上注册全局事件
	e.RegisterWithPriority("OnSkill", GlobalEvent, 900)

	// 调用事件，所有注册的同名函数都会被调用
	e.Call("OnSkill", 100)
}
