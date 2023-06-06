package balance

import (
	"errors"
	"math/rand"
)

// RandomBalance 随机负载均衡
type RandomBalance struct {
	curIndex int

	rss []string
}

func (r *RandomBalance) Add(params ...string) error {
	if len(params) == 0 {
		return errors.New("params len 1 at least")
	}
	r.rss = append(r.rss, params...)

	return nil
}

func (r *RandomBalance) Next() string {
	if len(r.rss) == 0 {
		return ""
	}
	if len(r.rss) == 1 {
		return r.rss[0]
	}
	r.curIndex = rand.Intn(len(r.rss))
	return r.rss[r.curIndex]
}

func (r *RandomBalance) Get() string {
	return r.Next()
}
