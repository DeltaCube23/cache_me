package cache_me

import "sync/atomic"

type Stats struct {
	hitCount  int32
	missCount int32
}

func newStats() *Stats {
	return &Stats{
		hitCount:  0,
		missCount: 0,
	}
}

func (st *Stats) incrementHit() {
	atomic.AddInt32(&st.hitCount, 1)
}

func (st *Stats) incrementMiss() {
	atomic.AddInt32(&st.missCount, 1)
}

func (st *Stats) HitCountFetch() int32 {
	return atomic.LoadInt32(&st.hitCount)
}

func (st *Stats) MissCountFetch() int32 {
	return atomic.LoadInt32(&st.missCount)
}

func (st *Stats) LookupCountFetch() int32 {
	return st.HitCountFetch() + st.MissCountFetch()
}
