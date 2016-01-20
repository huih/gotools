package fifolist
import (
	"sync"
	"container/list"
)

type FIFOList struct {
	list *list.List
	mutex sync.Mutex
}

func (l *FIFOList) Init() *FIFOList {
	l.list = list.New()
	return l
}

func (l *FIFOList) Add(value interface{}) {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	l.list.PushFront(value)
}

func (l *FIFOList) Get() interface {} {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	v := l.list.Back()
	
	if v == nil {
		return nil
	}
	return v.Value
}

func (l *FIFOList) GetAndRemove() interface {} {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	v := l.list.Back()
	
	if v == nil {
		return nil
	}
	
	l.list.Remove(v)
	
	return v.Value
}

func (l *FIFOList) Length() int {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	
	return l.list.Len()
}

func New() *FIFOList { return new(FIFOList).Init() }
