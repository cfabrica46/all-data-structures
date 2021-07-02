package queue

import (
	"errors"
	"fmt"
	"time"
)

type Job struct {
	next        *Job
	ID          int
	Delay       time.Duration
	JobFunction func()
}

type JobQueue struct {
	length, maxJobs int
	head, tail      *Job
}

var JQ *JobQueue

var idAux int

func NewJobQueue() *JobQueue {
	JQ = &JobQueue{maxJobs: 5}

	go func() {
		for {
			time.Sleep(time.Second)
			if !JQ.IsEmpty() {
				time.Sleep(JQ.head.Delay)
				JQ.head.JobFunction()
				j := outToQueue()
				fmt.Printf("Has been Completed Job ID:%v\n", j.ID)
			} else {
				fmt.Println("The queue is empty")
			}
		}
	}()
	return JQ
}

func NewJob(delay time.Duration, jFunc func()) *Job {
	idAux++
	return &Job{ID: idAux, Delay: time.Second, JobFunction: jFunc}
}

func (q JobQueue) String() string {
	var b string
	for !q.IsEmpty() {
		b += fmt.Sprintf("%v -> ", q.head.ID)
		q.head = q.head.next
	}
	return b
}

func (q JobQueue) Len() int {
	return q.length
}

func (q JobQueue) IsEmpty() (empty bool) {
	if q.head == nil {
		empty = true
	}
	return
}

func (q *JobQueue) JoinIntoQueue(j *Job) (err error) {
	if q.length == q.maxJobs {
		err = errors.New("the queue is full")
		return
	}

	if q.IsEmpty() {
		q.head = j
		q.tail = j
		q.length++
	} else {
		q.tail.next = j
		q.tail = j
		q.length++
	}
	return
}

func outToQueue() (j *Job) {

	j = JQ.head
	JQ.head = JQ.head.next
	if JQ.IsEmpty() {
		JQ.tail = nil
	}
	JQ.length--

	return
}
