package zapWrapper

import (
	"sync"
	"testing"
)

func TestLogger(t *testing.T) {
	Init("./test.logger", "info")

	wg := sync.WaitGroup{}

	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			l1 := NewSugar("hello")

			for i := 0; i < 1000; i++ {

				l1.Info("this is a test")
			}

		}()
	}

	wg.Wait()
}

// 根据日志级别将日志输出到不同的文件
func TestOutputToDifferentFileAccordingToLevel(t *testing.T) {
	Init2(&LoggerConfig{InfoPath: "./info.log", DefaultPath: "./debug.log"})
	logger := NewSugar("test")
	logger.Debug("hello")
	logger.Info("world")

}
