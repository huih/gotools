package fixedlist

import (
	"sync"
	"container/list"
)

type FixedList struct {
	list *list.List
	mutex sync.Mutex
	fixedLen int
}

func (self *FixedList) Init(list_size int) *FixedList {
	self.fixedLen = list_size
	self.list = list.New()
	return self
}

func (self *FixedList) Add(value interface{}) (interface{}) {
	self.mutex.Lock()
	defer self.mutex.Unlock()
	
	if self.list.Len() >= self.fixedLen {
		v := self.list.Front()
		self.list.Remove(v)
		self.list.PushBack(value)
		
		return v.Value
	}
	
	self.list.PushBack(value)

	return nil
}

func (self *FixedList) PopFront() (interface{}) {
	self.mutex.Lock()
	defer self.mutex.Unlock()
	
	v := self.list.Front()
	if v == nil {
		return nil
	}
	self.list.Remove(v)
	
	return v.Value
}

func New(list_size int) *FixedList { 
	return new(FixedList).Init(list_size) 
}