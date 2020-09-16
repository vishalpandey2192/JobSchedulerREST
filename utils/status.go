package utils

type Status string

const(
	QUEUED      Status = "QUEUED"
	IN_PROGRESS        = "IN_PROGRESS"
	CONCLUDED          = "CONCLUDED"
)
