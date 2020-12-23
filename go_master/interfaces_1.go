package main

import "log"

type User struct {
	Name  string
	Email string
}

type org struct {
	Name  string
	Email string
}

type Notifier interface {
	Notify() error
}

func (u User) Notify() error {
	log.Printf("User: Sending User Email To %s<%s>\n",
		u.Name,
		u.Email)

	return nil
}

func (o org) Notify() error {
	log.Printf("User: Sending Orginization Email To %s<%s>\n",
		o.Name,
		o.Email)

	return nil
}

func SendNotification(notify Notifier) error {
	return notify.Notify()
}

func main() {
	// Value of type User can be used to call the method
	// with a value receiver.
	bill := User{"Bill", "bill@email.com"}
	SendNotification(bill)

	// Pointer of type User can also be used to call a method
	// with a value receiver.
	jill := &User{"Jill", "jill@email.com"}
	SendNotification(jill)

	myorg := org{"Bhuvan", "venk@email.com"}
	SendNotification(myorg)
}
