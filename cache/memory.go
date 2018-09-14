package cache

import (
	"fmt"
	"sync"
	"time"
)

type mem struct {
	value    interface{}
	expireAt time.Time
}

// Memory ...
type Memory struct {
	sync.Mutex
	items map[string]mem
	cap   int
}

// NewMemory ...
func NewMemory() *Memory {
	return &Memory{
		items: map[string]mem{},
	}
}

func (r *Memory) memClean() {
	l := len(r.items)
	if float64(r.cap-l)/float64(l) > 0.4 {
		fmt.Println("clean memory cache")
		r.Lock()
		defer r.Unlock()
		// 有0.4的数据没有使用，清理一下

		items := make(map[string]mem, len(r.items))
		for k, v := range r.items {
			items[k] = mem{value: v.value, expireAt: v.expireAt}
		}
		r.items = items
		r.cap = len(items)
	}
}

// Get ...
func (r *Memory) Get(key string) interface{} {
	r.memClean()

	item, ok := r.items[key]
	if !ok {
		return nil
	}

	if item.expireAt.Sub(time.Now()) <= 0 {
		delete(r.items, key)
		return nil
	}

	return item.value
}

// Set ...
func (r *Memory) Set(key string, val interface{}, timeout time.Duration) error {
	r.memClean()

	if _, ok := r.items[key]; !ok {
		r.cap++
	}

	r.items[key] = mem{
		value:    val,
		expireAt: time.Now().Add(timeout),
	}

	return nil
}

// IsExist ...
func (r *Memory) IsExist(key string) bool {
	r.memClean()

	item, ok := r.items[key]
	if !ok {
		return false
	}

	if item.expireAt.Sub(time.Now()) <= 0 {
		delete(r.items, key)
		return false
	}
	return true
}

// Delete ...
func (r *Memory) Delete(key string) error {
	r.memClean()

	delete(r.items, key)
	return nil
}
