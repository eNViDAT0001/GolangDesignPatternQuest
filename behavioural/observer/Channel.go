package observer

type Channel interface {
	Register(observer Observer)
	DeRegister(observer Observer)
	NotifyAll()
}
type channel struct {
	users []Observer
	name  string
}

func (c *channel) Register(observer Observer) {
	c.users = append(c.users, observer)
}
func (c *channel) DeRegister(observer Observer) {
	c.users = removeFromSlice(c.users, observer)
}
func (c *channel) NotifyAll() {
	for _, user := range c.users {
		user.Update(c.name)
	}
}
func removeFromSlice(observerList []Observer, observerToRemove Observer) []Observer {
	finalIndex := len(observerList) - 1
	for i, user := range observerList {
		if user.GetID() == observerToRemove.GetID() {
			observerList[finalIndex], observerList[i] = observerList[i], observerList[finalIndex]
			return observerList[:finalIndex]
		}
	}
	return observerList
}

func CreateChannel(name string) Channel {
	return &channel{name: name}
}
