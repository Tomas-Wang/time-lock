# time-lock
local lock with timeout release

# Usage
## spinlock
```
    tl := NewTimedLock(5 * time.Second)
    tl.Lock()
```
## none spinlock
```
    tl := NewTimedLock(5 * time.Second)
    err := tl.TryLock()
    if err != nil {
        fmt.Println("lock failed")
    }
```


