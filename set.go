package main

import (
	"sync"
)

// Set 存放客户端请求字符串
type Set struct {
	values map[interface{}]struct{}
	mu sync.RWMutex
}

var fillValue = struct{}{}

// NewSet 实列化一个set
func NewSet() *Set {
	return &Set{ values: make(map[interface{}]struct{}) }
}

// Add 新增
func (s *Set) Add(value interface{}) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.values[value] = fillValue
}

// Has 判断给定值是否存在
func (s *Set) Has(value interface{}) bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	_, ok := s.values[value]
	return ok
}

// Clear 清空Set
func (s *Set) Clear() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.values = make(map[interface{}]struct{})
}

// Len 获取数据条数
func (s *Set) Len() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return len(s.values)
}
