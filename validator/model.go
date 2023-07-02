package validator

import (
	"fmt"
	"sort"
	"time"
)

type ProxyResult struct {
	Proxy    string
	Latency  time.Duration
	IsOnline bool
	Result   string
}

type ProxiesResult []ProxyResult

func (p ProxiesResult) SortByLatency() ProxiesResult {
	sort.Slice(p, func(i, j int) bool {
		return p[i].Latency < p[j].Latency
	})

	return p
}

func (ps ProxiesResult) FilterOnline() ProxiesResult {
	newPs := ProxiesResult{}
	for _, p := range ps {
		if p.IsOnline {
			newPs = append(newPs, p)
		}
	}
	return newPs
}

func (ps ProxiesResult) Print() {
	fmt.Printf("Proxy\t\t\t\tLatency \tStatus \t Result\n")
	for _, p := range ps {
		fmt.Printf("%s  \t%s\t%t\t %s\n", p.Proxy, p.Latency, p.IsOnline, p.Result)
	}
}
