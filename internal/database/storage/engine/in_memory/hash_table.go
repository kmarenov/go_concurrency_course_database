package in_memory

var HashTableBuilder = func() hashTable {
	return NewHashTable()
}

type HashTable struct {
	data map[string]string
}

func NewHashTable() *HashTable {
	return &HashTable{
		data: make(map[string]string),
	}
}

func (s *HashTable) Set(key, value string) {
	s.data[key] = value
}

func (s *HashTable) Get(key string) (string, bool) {
	value, found := s.data[key]
	return value, found
}

func (s *HashTable) Del(key string) {
	delete(s.data, key)
}
