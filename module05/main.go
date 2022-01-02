package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

type Bucket struct {
	sync.RWMutex
	TotalReq  int
	FailedReq int
	Timestamp time.Time
}

func NewBucket() *Bucket {
	return &Bucket{
		Timestamp: time.Now(),
	}
}

func (b *Bucket) Record(result bool) {
	b.Lock()
	defer b.Unlock()
	if !result {
		b.FailedReq++
	}
	b.TotalReq++
}

type RollingWindow struct {
	sync.RWMutex
	broken bool
	// 滑动窗口大小
	size int
	// 桶队列
	buckets []*Bucket
	// 触发熔断的请求总数阈值
	reqThreshold int
	// 出发熔断的失败率阈值
	failedThreshold float64
	// 上次熔断发生时间
	lastBreakTime time.Time
	seeker        bool
	// 熔断恢复的时间间隔
	brokeTimeGap time.Duration
}

// 新建滑动窗口
func NewRollingWindow(
	size int,
	reqThreshold int,
	failedThreshold float64,
	brokeTimeGap time.Duration,
) *RollingWindow {
	return &RollingWindow{
		size:            size,
		buckets:         make([]*Bucket, 0, size),
		reqThreshold:    reqThreshold,
		failedThreshold: failedThreshold,
		brokeTimeGap:    brokeTimeGap,
	}
}

func (r *RollingWindow) AppendBucket() {
	r.Lock()
	defer r.Unlock()
	r.buckets = append(r.buckets, NewBucket())
	if !(len(r.buckets) < r.size+1) {
		r.buckets = r.buckets[1:]
	}
}

func (r *RollingWindow) GetBucket() *Bucket {
	if len(r.buckets) == 0 {
		r.AppendBucket()
	}
	return r.buckets[len(r.buckets)-1]
}

func (r *RollingWindow) RecordReqResult(result bool) {
	r.GetBucket().Record(result)
}

func (r *RollingWindow) ShowAllBucket() {
	for _, v := range r.buckets {
		fmt.Printf("id: [%v] | total: [%d] | failed: [%d]\n", v.Timestamp, v.TotalReq, v.FailedReq)
	}
}

func (r *RollingWindow) Launch() {
	go func() {
		for {
			r.AppendBucket()
			time.Sleep(time.Millisecond * 100)
		}
	}()
}

func (r *RollingWindow) BreakJudgement() bool {
	r.RLock()
	defer r.RUnlock()
	total := 0
	failed := 0
	for _, v := range r.buckets {
		total += v.TotalReq
		failed += v.FailedReq
	}
	if float64(failed)/float64(total) > r.failedThreshold && total > r.reqThreshold {
		return true
	}
	return false
}

func (r *RollingWindow) Monitor() {
	go func() {
		for {
			if r.broken {
				if r.OverBrokenTimeGap() {
					r.Lock()
					r.broken = false
					r.Unlock()
				}
				continue
			}
			if r.BreakJudgement() {
				r.Lock()
				r.broken = true
				r.lastBreakTime = time.Now()
				r.Unlock()
			}
		}
	}()
}

func (r *RollingWindow) OverBrokenTimeGap() bool {
	return time.Since(r.lastBreakTime) > r.brokeTimeGap
}

func (r *RollingWindow) ShowStatus() {
	go func() {
		for {
			log.Println(r.broken)
			time.Sleep(time.Second)
		}
	}()
}

func (r *RollingWindow) Broken() bool {
	return r.broken
}

func (r *RollingWindow) SetSeeker(status bool) {
	r.Lock()
	defer r.Unlock()
}

func (r *RollingWindow) Seeker() bool {
	return r.seeker
}

func main() {
	r := NewRollingWindow(10,
		50,
		0.8,
		time.Second*5)
	r.Launch()
	r.Monitor()
	r.ShowStatus()
	select {}
}
