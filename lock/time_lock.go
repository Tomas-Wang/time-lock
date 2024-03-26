package lock

import (
	"sync"
	"time"
)

// TimedLock 添加时长的本地锁
type TimedLock struct {
	mu      sync.RWMutex
	timeout time.Duration
	stopCh  chan struct{}
}

func NewTimedLock(timeout time.Duration) *TimedLock {
	return &TimedLock{
		timeout: timeout,
		stopCh:  make(chan struct{}),
	}
}

func (tl *TimedLock) Lock() {
	tl.mu.Lock()

	// 启动一个goroutine来监听超时信号
	go func() {
		select {
		case <-time.After(tl.timeout):
			// 超时后解锁
			tl.mu.Unlock()
		case <-tl.stopCh:
			// 正常解锁时关闭通道，防止goroutine泄露
			return
		}
	}()
}

func (tl *TimedLock) Unlock() {
	// 发送信号给监听超时的goroutine，防止其解锁
	close(tl.stopCh)
	tl.mu.Unlock()
}
