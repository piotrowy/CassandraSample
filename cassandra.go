package main

import (
	"github.com/gocql/gocql"
	"time"
	"sync"
)

var (
	cluster *gocql.ClusterConfig
	WP WorkerPool
)

type (
	QueryFn func(session *gocql.Session)

	WorkerLauncher interface {
		Launch(jobs <-chan QueryFn, sessionPool chan *gocql.Session)
	}

	WorkerPool interface {
		LaunchWorker()
		RemoveWorker()
		DoJob(fn QueryFn)
		Close()
	}

	Worker interface {
		WorkerLauncher
		Stop()
	}

	workerPool struct {
		sessions chan *gocql.Session
		jobs     chan QueryFn
		workers  []*Worker
	}

	worker struct {
		close chan bool
	}
)

func NewWorkerPoll(workersNum, sessionNum, jobsNum int, cluster gocql.ClusterConfig) WorkerPool {
	var (
		sessionsCh            = make(chan *gocql.Session, sessionNum)
		jobsCh                = make(chan QueryFn, jobsNum)
		wp         WorkerPool = &workerPool{sessions: sessionsCh, jobs: jobsCh}
		wg                    = sync.WaitGroup{}
	)
	wg.Add(2)

	go func(ch chan<- *gocql.Session, wg sync.WaitGroup) {
		for i := 0; i < sessionNum; {
			session, err := cluster.CreateSession()
			if err != nil {
				continue
			}
			sessionsCh <- session
			i++
		}
		wg.Done()
	}(sessionsCh, wg)

	go func(wg sync.WaitGroup) {
		for i := 0; i < workersNum; i++ {
			wp.LaunchWorker()
		}
		wg.Done()
	}(wg)

	wg.Wait()
	return wp
}

func (wp *workerPool) DoJob(fn QueryFn) {
	for {
		select {
		case wp.jobs <- fn:
			return
		case <-time.After(time.Second * 5):
			wp.LaunchWorker()
		}
	}
}

func (wp *workerPool) Close() {
	for _, v := range wp.workers {
		(*v).Stop()
	}
	close(wp.sessions)
	for s := range wp.sessions {
		s.Close()
	}
}

func (wp *workerPool) LaunchWorker() {
	var w Worker = &worker{
		close: make(chan bool),
	}
	w.Launch(wp.jobs, wp.sessions)
	wp.workers = append(wp.workers, &w)
}

func (wp *workerPool) RemoveWorker() {
	(*wp.workers[0]).Stop()
	wp.workers = wp.workers[1:]
}

func (w *worker) Stop() {
	w.close <- true
}

func (w *worker) Launch(jobs <-chan QueryFn, sessionPool chan *gocql.Session) {
	go func() {
		for {
			select {
			case fn := <-jobs:
				s := <-sessionPool
				fn(s)
				sessionPool <- s
			case <-w.close:
				return
			}
		}
	}()
}

func initCassandra(ips ...string) {
	cluster = gocql.NewCluster(ips...)
	cluster.Keyspace = "Pub"
	cluster.Consistency = gocql.Quorum
}
