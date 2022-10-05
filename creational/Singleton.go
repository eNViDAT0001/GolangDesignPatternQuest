package creational

import "sync"

type Counter interface {
	AddOne() int
	AddNum(int) int
}
type counter struct {
	count int
}

var instance *counter

func (c *counter) AddOne() int {
	c.count++
	return c.count
}
func (c *counter) AddNum(n int) int {
	c.count += n
	return c.count
}

func GetInstance() Counter {
	once := sync.Once{}
	once.Do(func() {
		instance = &counter{count: 0}
	})
	return instance
}
