package cset

import (
	"constraints"
	"sync"
)

type Set[K constraints.Ordered] interface {
	Clear() Set[K]
	Add(a ...K) Set[K]
	Delete(a ...K) Set[K]
	Immutable() ImmutableSet[K]

	Contains(a ...K) bool
	Keys() []K
	Range(apply func(a K))
}

type ImmutableSet[K constraints.Ordered] interface {
	Contains(a ...K) bool
	Keys() []K
	Range(apply func(a K))
}

type set[K constraints.Ordered] struct {
	immutable bool
	mu        sync.RWMutex
	content   map[K]bool
}

func New[K constraints.Ordered]() Set[K] {
	s := &set[K]{
		content: map[K]bool{},
	}
	return s
}

func From[K constraints.Ordered](keys ...K) Set[K] {
	return New[K]().Add(keys...)
}

func (s *set[K]) Immutable() ImmutableSet[K] {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.immutable = true
	return s
}

func (s *set[K]) Clear() Set[K] {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.immutable {
		panic("can not clear an immutable set")
	}

	s.content = make(map[K]bool)
	return s
}

func (s *set[K]) Add(all ...K) Set[K] {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.immutable {
		panic("can not add to an immutable set")
	}
	for _, a := range all {
		s.content[a] = true
	}
	return s
}

func (s *set[K]) Delete(all ...K) Set[K] {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.immutable {
		panic("can not remove from an immutable set")
	}
	for _, a := range all {
		delete(s.content, a)
	}
	return s
}

func (s *set[K]) Contains(all ...K) bool {
	s.mu.RLock()
	defer s.mu.RUnlock()

	for _, a := range all {
		if !s.content[a] {
			return false
		}
	}
	return true
}

func (s *set[K]) Keys() []K {
	s.mu.RLock()
	defer s.mu.RUnlock()
	var keys []K
	for k, _ := range s.content {
		keys = append(keys, k)
	}
	return keys
}

func (s *set[K]) Range(apply func(a K)) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	for k, _ := range s.content {
		apply(k)
	}
}
func (s *set[K]) Len() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return len(s.content)
}