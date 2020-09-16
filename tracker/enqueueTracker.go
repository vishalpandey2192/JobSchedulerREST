package tracker

import (
	"errors"
	"go-rest-api-develop/manager/job"
	"sync"
)

type EnqueueManager struct {
	state map[int]*job.Job
	availableJobs []int
	sync.RWMutex
}

func NewEnqueueManager() *EnqueueManager {
	return &EnqueueManager{
		state: make(map[int]*job.Job),
		availableJobs: []int{},
	}
}

func (manager *EnqueueManager) Contains(id int) *job.Job {

	// acquire read lock
	manager.RLock()
	defer manager.RUnlock()
	if val, ok := manager.state[id]; ok {
		return val
	}
	return nil
}

func (manager *EnqueueManager) Add(job *job.Job) (int, error) {

	// acquire read lock
	manager.Lock()
	defer manager.Unlock()

	if _, ok := manager.state[job.ID]; ok {
		return -1, errors.New("already queued")
	}

	// set type
	job.IsQueued()

	// add to available jobs
	manager.availableJobs = append(manager.availableJobs, job.ID)

	manager.state[job.ID] = job
	return job.ID, nil
}

func (manager *EnqueueManager) Remove(id int) (*job.Job, error) {

	// acquire read lock
	manager.Lock()
	defer manager.Unlock()

	if val, ok := manager.state[id]; ok {
		delete(manager.state, id)
		return val, nil
	}

	return nil, errors.New("job not found")
}


func (manager *EnqueueManager) GetAvailableJob() (int, error) {

	// acquire read lock
	manager.Lock()
	defer manager.Unlock()

	if len(manager.availableJobs) > 0 {
		val := manager.availableJobs[0]
		manager.availableJobs = manager.availableJobs[1:]
		return val, nil
	}

	return -1, errors.New("no jobs available")
}



