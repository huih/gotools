package fifolist_test

import (
	"github.com/gotools/lists"
	"github.com/gotools/logs"
)

func ExampleFIFOList(){
	fifoL := fifolist.New()
	fifoL.Add(1)
	fifoL.Add(2)
	fifoL.Add(3)
	fifoL.Add(4)
	
	v := fifoL.GetAndRemove()
	logs.Debug("v: %d", v)
	v = fifoL.GetAndRemove()
	logs.Debug("v: %d", v)
	v = fifoL.GetAndRemove()
	logs.Debug("v: %d", v)
	v = fifoL.GetAndRemove()
	logs.Debug("v: %d", v)
}