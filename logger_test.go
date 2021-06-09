package zapWrapper

import (
	"sync"
	"testing"
)

func TestLogger(t *testing.T) {
	Init("./test.log", "debug")

	wg := sync.WaitGroup{}

	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			l1 := NewSugar("hello")

			for i := 0; i < 1000; i++ {

				l1.Debug("this is a test")
			}

		}()
	}

	wg.Wait()
}

// 根据日志级别将日志输出到不同的文件
func TestOutputToDifferentFileAccordingToLevel(t *testing.T) {
	Init2(&LoggerConfig{InfoPath: "./log/info.log", DebugPath: "./log/debug.log"})
	logger := NewSugar("test")
	logger.Debug("hello")
	logger.Info("world")
	logger.Error("error")
}
