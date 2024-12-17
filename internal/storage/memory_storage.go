package storage

import (
	"errors"
	"sort"
	"sync"
)

type Storage interface {
	GetPackSizes() []int
	AddPackSize(size int)
	RemovePackSize(size int) error
}

type memoryStorage struct {
	mu        sync.Mutex
	packSizes []int
}

func NewMemoryStorage(initial []int) Storage {
	s := &memoryStorage{
		packSizes: append([]int(nil), initial...),
	}
	sort.Ints(s.packSizes)
	return s
}

func (m *memoryStorage) GetPackSizes() []int {
	m.mu.Lock()
	defer m.mu.Unlock()
	cp := make([]int, len(m.packSizes))
	copy(cp, m.packSizes)
	return cp
}

func (m *memoryStorage) AddPackSize(size int) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.packSizes = append(m.packSizes, size)
	sort.Ints(m.packSizes)
}

func (m *memoryStorage) RemovePackSize(size int) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	index := -1
	for i, p := range m.packSizes {
		if p == size {
			index = i
			break
		}
	}
	if index == -1 {
		return errors.New("not found")
	}

	m.packSizes = append(m.packSizes[:index], m.packSizes[index+1:]...)
	return nil
}
