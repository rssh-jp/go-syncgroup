package syncgroup

import (
	"sync"
)

// SyncGroup is combines sync.WaitGroup and semaphore
type SyncGroup struct {
	wg        sync.WaitGroup
	semaphore chan struct{}
}

// New is create SyncGroup instance
func New(semaphoreCount int) *SyncGroup {
	return &SyncGroup{
		semaphore: make(chan struct{}, semaphoreCount),
	}
}

// Add provides functionality similar to WaitGroup.Add
func (s *SyncGroup) Add() {
	s.wg.Add(1)
	s.semaphore <- struct{}{}
}

// Done provides functionality similar to WaitGroup.Done
func (s *SyncGroup) Done() {
	s.wg.Done()
	<-s.semaphore
}

// Wait provides functionality similar to WaitGroup.Wait
func (s *SyncGroup) Wait() {
	s.wg.Wait()
}

// Close is close instance
func (s *SyncGroup) Close() {
	close(s.semaphore)
}
