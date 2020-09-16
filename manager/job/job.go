package job

import "go-rest-api-develop/utils"

type Job struct {
	ID int `json:"ID"`
	Type   utils.Type `json:"Type"`
	Status utils.Status      `json:"Status"`
}


func (j *Job) IsQueued() {
	j.Status = utils.QUEUED
}

func (j *Job) IsInProgress() {
	j.Status = utils.IN_PROGRESS
}

func (j *Job) IsConcluded() {
	j.Status = utils.CONCLUDED
}
