package event

import "fmt"

type event struct {
	priority int
	function func(interface{})
}

type EventService struct {
	handle map[string][]event
}

func New() *EventService {
	return &EventService{
		// 实例化一个通过字符串映射函数切片的map
		handle: make(map[string][]event),
	}
}

// 调用事件
func (e *EventService) Call(name string, param interface{}) {
	// 通过名字找到事件列表
	list := e.handle[name]
	fmt.Println("list:", list)

	// 遍历这个事件的所有回调
	for _, event := range list {
		// 传入参数调用回调
		event.function(param)
	}
}

// 注册事件，提供事件名和回调函数
func (e *EventService) Register(name string, callback func(interface{})) {
	e.RegisterWithPriority(name, callback, 50)
}

// 注册事件，提供事件名和回调函数
func (e *EventService) RegisterWithPriority(name string, callback func(interface{}), priority int) {

	// 通过名字查找事件列表
	list := e.handle[name]

	// 在列表切片中添加函数
	list = append(list, event{
		priority: priority,
		function: callback,
	})

	// 将修改的事件列表切片保存回去
	e.handle[name] = list
}
