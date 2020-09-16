package manager

import (
	"errors"
	"go-rest-api-develop/manager/job"
	"go-rest-api-develop/tracker"

)

type JobManager struct {
	enqueueTracker *tracker.EnqueueManager
	dequeueTracker *tracker.DequeueManager
	completedJobs *tracker.CompletedManager
}

func (jm *JobManager) Conclude(id int) (*job.Job,  error)  {

	// check if in dequeued jobs
	if job := jm.dequeueTracker.Contains(id); job != nil {

		// add to completed
		jm.completedJobs.Add(job)

		// remove from completed
		jm.dequeueTracker.Remove(id)
		return job, nil

	}

	return nil, errors.New("job not dequeued yet")

}

func (jm *JobManager) Enqueue(job *job.Job) {
	// add to enqueueTracker
	jm.enqueueTracker.Add(job)
}

func (jm *JobManager) Dequeue() (*job.Job, error){

	var job *job.Job

	id , err := jm.enqueueTracker.GetAvailableJob()
	if err != nil {
		return nil, err
	}

	// remove from enqueue
	job, err = jm.enqueueTracker.Remove(id)
	if err != nil {
		return nil, err
	}

	job.IsInProgress()

	jm.dequeueTracker.Add(job)

	return job, nil
	}

func (jm *JobManager) Contains(id int) *job.Job {
	// remove job from enqueue
	// add to dequeue

	if job := jm.enqueueTracker.Contains(id); job != nil {
		return job
	}

	if job := jm.dequeueTracker.Contains(id); job != nil {
		return job
	}

	if job := jm.completedJobs.Contains(id); job != nil {
		return job
	}

	return nil
}



func NewJobManager() *JobManager {
	return &JobManager{
		enqueueTracker: tracker.NewEnqueueManager(),
		dequeueTracker: tracker.NewDequeueManager(),
		completedJobs: tracker.NewCompletedManager(),
	}
}
