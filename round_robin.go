package balance

import "errors"

// RoundRobinBalance 轮询负载均衡
type RoundRobinBalance struct {
	curIndex int
	rss      []string
}

func (r *RoundRobinBalance) Add(params ...string) error {
	if len(params) == 0 {
		return errors.New("params len 1 at least")
	}

	r.rss = append(r.rss, params...)
	return nil
}

func (r *RoundRobinBalance) Next() string {
	lens := len(r.rss)
	if lens == 0 {
		return ""
	}

	if r.curIndex >= lens {
		r.curIndex = 0
	}

	curAddr := r.rss[r.curIndex]
	r.curIndex = (r.curIndex + 1) % lens
	return curAddr
}

func (r *RoundRobinBalance) Get() string {
	return r.Next()
}
