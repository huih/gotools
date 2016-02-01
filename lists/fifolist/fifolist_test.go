package fifolist
import (
	"testing"
	"github.com/gotools/logs"
)

func BenchmarkUseFifoList(b *testing.B) {
	fifoL := New()
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