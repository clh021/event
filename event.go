package event

type Event struct {
	eventMap map[string][]func(interface{})
}

func New() *Event {
	return &Event{
		// 实例化一个通过字符串映射函数切片的map
		eventMap: make(map[string][]func(interface{})),
	}
}

// 调用事件
func (e *Event) Call(name string, param interface{}) {
	// 通过名字找到事件列表
	list := e.eventMap[name]

	// 遍历这个事件的所有回调
	for _, callback := range list {
		// 传入参数调用回调
		callback(param)
	}
}

// 注册事件，提供事件名和回调函数
func (e *Event) Register(name string, callback func(interface{})) {

	// 通过名字查找事件列表
	list := e.eventMap[name]

	// 在列表切片中添加函数
	list = append(list, callback)

	// 将修改的事件列表切片保存回去
	e.eventMap[name] = list
}
