package manager_test

import (
	"github.com/stretchr/testify/assert"
	"go-rest-api-develop/manager"
	job2 "go-rest-api-develop/manager/job"
	"testing"
)


func TestNewJobManager(t *testing.T) {
	got := manager.NewJobManager()
	assert.NotNil(t, got)
}

func TestEnqueue(t *testing.T) {
	mgr := manager.NewJobManager()
	assert.NotNil(t, mgr)
	job := &job2.Job{
		ID:     1,
		Type:   "test",
		Status: "test",
	}
	mgr.Enqueue(job)

	assert.NotNil(t, mgr.Contains(1))
}

func TestDequeue(t *testing.T) {
	mgr := manager.NewJobManager()
	assert.NotNil(t, mgr)
	job := &job2.Job{
		ID:     1,
		Type:   "test",
		Status: "test",
	}
	mgr.Enqueue(job)
	assert.NotNil(t, mgr.Contains(1))

	mgr.Dequeue()
	assert.NotNil(t, mgr.Contains(1))
}