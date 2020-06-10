package main

import (
	"time"
	"fmt"
)


type  LB struct {
	workers  chan  *worker
	jobC     chan func()
}

func  newLB(workers ...int) LB  {
	lb := LB{
		workers:make(chan *worker,len(workers)),
		jobC:make(chan func()),
	}


	for i:=0;i< len(workers);i++{
		worker := newWorker(workers[i])
		worker.lb = lb
		lb.workers <- worker
	}

	go func() {
		for {
			f := <-lb.jobC
			avaliableW := <-lb.workers
			avaliableW.submit(f)
		}
	}()

	return lb
}

func (lb LB) submit(f func())  {
	go func() {
		fmt.Println("lb.submit.begin")
		lb.jobC <- f
		fmt.Println("lb.submit.end")
	}()
}


type worker struct {
	channel chan struct{}
	lb LB
}

func newWorker(parallel int)  *worker {
	return  &worker{channel:make(chan struct{},parallel)}
}

func (w *worker)submit(f func())  {
	select {
	case w.channel <- struct{}{}:
		go func() {
			f();
			<- w.channel
			w.lb.workers <- w
		}()
	default:
		fmt.Println("worker.submit.defalut")
	}
}

func main()  {
	lb := newLB(1,1,1,2)
	for i:=0;i<10;i++{
		//j  := i
		f := func() {
			time.Sleep(time.Second * 1)
			fmt.Println("doing  work",i)
		}
		lb.submit(f)
	}

	time.Sleep(time.Second * 10)

}