package tracker

import (
	"errors"
	"go-rest-api-develop/manager/job"
	"sync"
)

type CompletedManager struct {
	state map[int]*job.Job
	sync.RWMutex
}

func NewCompletedManager() *CompletedManager {
	return &CompletedManager{
		state: make(map[int]*job.Job),
	}
}

func (manager *CompletedManager) Contains(id int) *job.Job {

	// acquire read lock
	manager.RLock()
	defer manager.RUnlock()
	if val, ok := manager.state[id]; ok {
		return val
	}
	return nil
}

func (manager *CompletedManager) Add(job *job.Job) (int, error) {

	// acquire read lock
	manager.Lock()
	defer manager.Unlock()

	if _, ok := manager.state[job.ID]; ok {
		return -1, errors.New("already in progress")
	}

	// set type
	job.IsConcluded()


	manager.state[job.ID] = job
	return job.ID, nil
}

func (manager *CompletedManager) Remove(id int) (*job.Job, error) {

	// acquire read lock
	manager.Lock()
	defer manager.Unlock()

	if val, ok := manager.state[id]; ok {
		delete(manager.state, id)
		return val, nil
	}

	return nil, errors.New("job not found")
}

