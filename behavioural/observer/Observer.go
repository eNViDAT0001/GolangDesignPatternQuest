package observer

type Observer interface {
	Update(interface{})
	GetID() string
}
