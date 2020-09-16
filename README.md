# Jobs-rest-api

## Prerequisites

* Have Go installed. If not, check it out [here](https://golang.org/doc/install)

* Also after installing, make sure you are working inside your GOPATH

### Description

* We are creating a JSON API that will allow users to enqueue, dequeue and conclude jobs


### How to Run
* run `go build -o queue.exe`
* run `./queue.exe`

This project runs on port 9000

### APIs

1. Enqueue Job 

    POST ``localhost:9000/jobs/enqueue``
    Payload ``{"ID":2,"Type":"TIME_CRITICAL","Status":"IN_PROGRESS"}``
    
    Response - 
    * ID if successful 200
    * Error if already present 400
    
2. Dequeue Job 

    GET ``localhost:9000/jobs/dequeue``
    
    Response - 
    * Return complete job struct if any job is available 
    * Error if no job available 400
    
3. Conclude Job 

    GET ``localhost:9000/jobs/{id}/conclude``
    
    Response - 
    * Return complete job struct if any job is available and was in progress
    * Error if job was not moved to in progress
    * Error if job not present
    
    
Developer - Vishal Pandey - vishalpandey92@outlook.com
    