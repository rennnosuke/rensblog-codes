package main

func main() {

	kvs := NewKeyValueStore()

	for i := 0; i < 10; i++ {
		go func(kvs *KeyValueStore) {
			kvs.set("key", "value")
			kvs.get("key")
		}(kvs)
	}

}

type KeyValueStore struct {
	m map[string]string
}

func NewKeyValueStore() *KeyValueStore {
	return &KeyValueStore{m: make(map[string]string)}
}

func (s *KeyValueStore) set(k, v string) {
	s.m[k] = v
}

func (s *KeyValueStore) get(k string) (string, bool) {
	v, ok := s.m[k]
	return v, ok
}
