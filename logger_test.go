package zapWrapper

import (
	"go.uber.org/zap/zapcore"
	"sync"
	"testing"
)

func TestLogger(t *testing.T) {
	Init("./test.log", zapcore.DebugLevel)

	wg := sync.WaitGroup{}

	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			l1 := NewSugar("hello")

			for i := 0; i < 1000; i++ {

				l1.Error("this is a test")
			}

		}()
	}

	wg.Wait()

}
