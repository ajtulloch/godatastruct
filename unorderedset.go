package godatastruct

import (
	"sync"
)

type Set interface {
	Set(key interface{}) bool
	Exists(key interface{}) bool
	Erase(key interface{}) bool
	Len() int
	Clear()
}

type unorderedSet struct {
	elementMap map[interface{}]bool
}

func NewUnorderedSet() Set {
	return &unorderedSet{make(map[interface{}]bool)}
}

func (s *unorderedSet) Exists(key interface{}) bool {
	_, exists := s.elementMap[key]
	return exists
}

func (s *unorderedSet) Set(key interface{}) bool {
	if s.Exists(key) {
		return false
	}

	s.elementMap[key] = true
	return true
}

func (s *unorderedSet) Erase(key interface{}) bool {
	existed := s.Exists(key)
	delete(s.elementMap, key)
	return existed
}

func (s *unorderedSet) Len() int {
	return len(s.elementMap)
}

func (s *unorderedSet) Clear() {
	for key := range s.elementMap {
		delete(s.elementMap, key)
	}
}

type threadSafeUnorderedSet struct {
	mutex        sync.RWMutex
	unorderedSet unorderedSet
}

func NewTheadSafeUnorderedSet() Set {
	s := new(threadSafeUnorderedSet)
	s.unorderedSet = unorderedSet{make(map[interface{}]bool)}
	return s
}

func (s *threadSafeUnorderedSet) Exists(key interface{}) bool {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	return s.unorderedSet.Exists(key)
}

func (s *threadSafeUnorderedSet) Set(key interface{}) bool {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	return s.unorderedSet.Set(key)
}

func (s *threadSafeUnorderedSet) Erase(key interface{}) bool {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	return s.unorderedSet.Erase(key)
}

func (s *threadSafeUnorderedSet) Len() int {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	return s.unorderedSet.Len()
}

func (s *threadSafeUnorderedSet) Clear() {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.unorderedSet.Clear()
}
