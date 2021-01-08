package main

import (
	"sync"
)

func main() {

	kvs := NewConcurrentKeyValueStore()

	for i := 0; i < 10; i++ {
		go func(kvs *ConcurrentKeyValueStore) {
			kvs.set("key", "value")
			kvs.get("key")
		}(kvs)
	}

}

type ConcurrentKeyValueStore struct {
	m  map[string]string
	mu sync.RWMutex
}

func NewConcurrentKeyValueStore() *ConcurrentKeyValueStore {
	return &ConcurrentKeyValueStore{m: make(map[string]string)}
}

func (s *ConcurrentKeyValueStore) set(k, v string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.m[k] = v
}

func (s *ConcurrentKeyValueStore) get(k string) (string, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	v, ok := s.m[k]
	return v, ok
}
