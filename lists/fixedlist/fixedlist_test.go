package fixedlist
import (
	"testing"
	"github.com/gotools/logs"
)

func BenchmarkFixedList(b *testing.B) {
	fixedL := New(2)
	fixedL.Add(1)
	fixedL.Add(2)
	fixedL.Add(3)
	fixedL.Add(4)
	
	for {
		v := fixedL.PopFront()
		if v == nil {
			break
		}
		
		logs.DebugS(v)
	}
}