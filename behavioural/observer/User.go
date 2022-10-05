package observer

import "fmt"

var id int = 0

type User interface {
	NewUser()
}
type user struct {
	ID       string
	username string
}

func (u *user) Update(itemName interface{}) {
	fmt.Printf("Notify to %s for %s\n", u.username, itemName)
}
func (u *user) GetID() string {
	return u.ID
}
func NewUser(name string) Observer {
	return &user{
		ID:       generateID(),
		username: name,
	}
}
func generateID() string {
	id++
	return string(id)
}
