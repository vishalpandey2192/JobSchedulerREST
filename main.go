package main

import (
	"encoding/json"
	"fmt"
	"go-rest-api-develop/manager/job"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"go-rest-api-develop/manager"
)

var tracker *manager.JobManager

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to my awesome job tracker!")
}

func enqueueHandler(w http.ResponseWriter, r *http.Request) {
	var newJob job.Job
	// Convert r.Body into a readable format
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Enter Valid Data")
	}

	json.Unmarshal(reqBody, &newJob)

	// check if Already Present
	if present := tracker.Contains(newJob.ID); present != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Job ID already present")
		return
	}

	// set status
	newJob.IsQueued()

	// Add the job to enqueue
	tracker.Enqueue(&newJob)

	// Return the 201 created status code
	w.WriteHeader(http.StatusCreated)
	// Return the newly created newJob
	json.NewEncoder(w).Encode(newJob.ID)
}

func concludeHandler(w http.ResponseWriter, r *http.Request) {
	// Get the ID from the url
	eventID, _ := strconv.Atoi(mux.Vars(r)["id"])

	// Get the details from an existing job
	// Use the blank identifier to avoid creating a value that will not be used
	job, err := tracker.Conclude(eventID)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(job)
	return
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(tracker.Contains(1))
}

func fetchHandler(w http.ResponseWriter, r *http.Request) {
	eventID, _ := strconv.Atoi(mux.Vars(r)["id"])

	job := tracker.Contains(eventID)

	if job != nil {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(job)
	}
}

func dequeueHandler(w http.ResponseWriter, r *http.Request) {

	id, err := tracker.Dequeue()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(id)
}


func main() {

	//create holder
	tracker = manager.NewJobManager()
	prefilledJob := &job.Job{
		ID:     1,
		Type:   "TIME_CRITICAL",
		Status: "IN_PROGRESS",
	}
	tracker.Enqueue(prefilledJob)



	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)


	api2 := router.PathPrefix("/jobs").Subrouter()
	api2.HandleFunc("/enqueue", enqueueHandler).Methods("POST")
	api2.HandleFunc("/sample", testHandler).Methods("GET")
	api2.HandleFunc("/dequeue", dequeueHandler).Methods("GET")
	api2.HandleFunc("/{id}/conclude", concludeHandler).Methods("GET")
	api2.HandleFunc("/{id}", fetchHandler).Methods("GET")



	log.Fatal(http.ListenAndServe(":9000", router))
}
