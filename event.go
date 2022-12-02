package event

import (
	"sort"
)

type event struct {
	priority int
	callback func(interface{})
}

type events []event

type EventService struct {
	handle map[string]events
}

func (e events) Len() int {
	return len(e)
}
func (e events) Less(i, j int) bool {
	eli := e[i]
	elj := e[j]
	// 原本应是小于的排在前面，但是这里业务需求希望数值越大越靠前
	return eli.priority > elj.priority
}
func (e events) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}

func New() *EventService {
	return &EventService{
		// 实例化一个通过字符串映射函数切片的map
		handle: make(map[string]events),
	}
}

// 调用事件
func (s *EventService) Call(name string, param interface{}) {
	// 通过名字找到事件列表
	list := s.handle[name]
	sort.Sort(events(list))

	// 遍历这个事件的所有回调
	for _, es := range list {
		// 传入参数调用回调
		es.callback(param)
	}
}

// 注册事件，提供事件名和回调函数，默认优先级
func (s *EventService) Register(name string, callback func(interface{})) {
	s.RegisterWithPriority(name, callback, 50)
}

// 注册事件，提供事件名和回调函数，支持自定义优先级
// 优先级数值越大优先级越高
func (s *EventService) RegisterWithPriority(name string, callback func(interface{}), priority int) {

	// 通过名字查找事件列表
	list := s.handle[name]

	// 在列表切片中添加函数
	list = append(list, event{
		priority: priority,
		callback: callback,
	})

	// 将修改的事件列表切片保存回去
	s.handle[name] = list
}
