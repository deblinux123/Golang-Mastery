package main

import (
	"fmt"
	"sync"
)

type safeMap struct {
	mu   sync.RWMutex
	data map[string]int
}

func NewSafeMap() *safeMap {
	return &safeMap{
		data: make(map[string]int),
	}
}

func (s *safeMap) Set(key string, value int) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.data[key] = value
}

func (s *safeMap) Get(key string) (int, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	value, ok := s.data[key]

	return value, ok
}

func main() {
	scores := make(map[string]int)

	var wg sync.WaitGroup
	var mu sync.Mutex

	for i := 0; i < 10; i++ {
		wg.Add(1)

		go func(id int) {
			defer wg.Done()

			mu.Lock()
			scores[fmt.Sprintf("user-%d", id)] = id
			mu.Unlock()
		}(i)
	}

	wg.Wait()
	fmt.Println(scores)

	newScore := NewSafeMap()

	newScore.Set("farid", 98)

	value, ok := newScore.Get("farid")

	if ok {
		fmt.Println(value)
	}
}
