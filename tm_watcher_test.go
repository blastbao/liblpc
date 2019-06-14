package liblpc

import (
	"fmt"
	"gitee.com/Puietel/std"
	"testing"
	"time"
)

func TestTMWatcher(t *testing.T) {
	loop, err := NewEventLoop()
	std.AssertError(err, "NewEventLoop")
	defer func() {
		std.AssertError(loop.Close(), "close loop")
	}()
	countDown := 5
	watcher, err := NewTimerWatcher(loop, ClockMonotonic, func(ins *Timer) {
		fmt.Println("ontick ", time.Now().String())
		countDown--
		if countDown == 0 {
			std.AssertError(ins.Stop(), "stop watcher")
			loop.Break()
			fmt.Println("timer stop")
		}
	})
	std.AssertError(err, "new timer watcher")
	defer func() { std.AssertError(watcher.Close(), "watcher close") }()
	watcher.Update(true)
	err = watcher.Start(1000, 1000)
	std.AssertError(err, "watcher start")
	loop.Run()
}
