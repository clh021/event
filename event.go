package event

import (
	"fmt"
	"sort"
	"strings"
)

type Event struct {
	Priority int
	Callback func(interface{}, interface{})
	Param    interface{}
}

type events []Event

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
	return eli.Priority > elj.Priority
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

// 根据触发事件名取出符合监听规则的事件函数列表
func (s *EventService) list(name string) events {
	var list events
	// handle map keys
	for kname, e := range s.handle {
		if strings.HasSuffix(kname, "*") {
			matchName := strings.TrimRight(kname, "*")
			if strings.HasPrefix(name, matchName) {
				list = append(list, e...)
			}
		} else if kname == name {
			list = append(list, e...)
		}
	}
	return list
}

func (s *EventService) Call(name string, callParam interface{}) {
	s.call(fmt.Sprintf("before.%s", name), callParam)
	s.call(name, callParam)
	s.call(fmt.Sprintf("after.%s", name), callParam)
}

// 调用事件
func (s *EventService) call(name string, callParam interface{}) {
	// 通过名字找到事件列表
	list := s.list(name)
	sort.Sort(events(list))

	// 遍历这个事件的所有回调
	for _, es := range list {
		// 传入参数调用回调
		es.Callback(es.Param, callParam)
	}
}

// 注册事件，提供事件名和回调函数，支持自定义优先级
// 优先级数值越大优先级越高
func (s *EventService) Register(name string, e Event) {

	// 通过名字查找事件列表
	list := s.handle[name]

	// 在列表切片中添加函数
	list = append(list, e)

	// 将修改的事件列表切片保存回去
	s.handle[name] = list
}
