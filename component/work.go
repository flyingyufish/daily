package component

import (
	"fmt"
)

/*
一主多从模式
*/
type Job struct {
	Num int
}

func NewJob(num int) Job {
	return Job{Num: num}
}

type Worker struct {
	WorkerChan chan chan Job
	JobChan    chan Job
	Result     map[int]int
	Quit       chan bool
}

func NewWorker(workChan chan chan Job, result map[int]int) *Worker {
	return &Worker{
		WorkerChan: workChan,
		JobChan:    make(chan Job, 1),
		Result:     result,
		Quit:       make(chan bool, 1),
	}
}

func (w *Worker) Start() {
	go func() {
		for {
			w.WorkerChan <- w.JobChan
			select {
			case job := <-w.JobChan:
				fmt.Println(job.Num)
				w.Result[job.Num] = job.Num * job.Num
			case q := <-w.Quit:
				if q {
					return
				}
			}
		}
	}()
}
func (w *Worker) Stop() {
	w.Quit <- true
}

var Count int = 2

type Master struct {
	WorkChan  chan chan Job
	QueueChan chan Job
	WorkList  []*Worker
	Result    map[int]int
}

func NewMaster(c int, result map[int]int) *Master {
	Count = c
	return &Master{
		WorkChan:  make(chan chan Job, Count),
		QueueChan: make(chan Job, 2*Count),
		Result:    result,
	}
}
func (m *Master) Run() {
	for i := 0; i < Count; i++ {
		work := NewWorker(m.WorkChan, m.Result)
		m.WorkList = append(m.WorkList, work)
		work.Start()
	}
	go m.dispth()
}
func (m *Master) dispth() {
	for {
		select {
		case job := <-m.QueueChan:
			j := <-m.WorkChan
			j <- job
		}
	}
}
func (m *Master) AddJob(num int) {
	j := Job{Num: num}
	m.QueueChan <- j
}
