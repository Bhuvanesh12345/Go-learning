package main

import "fmt"

type User struct {
	Name  string
	Email string
}

func (u *User) Notify() error {
	fmt.Printf("My name is %v. Please reach out to my email: %v\n", u.Name, u.Email)
	return nil
}

type Notifier interface {
	Notify() error
}

func main() {
	// Value of type User can be used to call the method
	// with a value receiver.
	bill := User{"Bill", "bill@email.com"}
	k := bill.Notify()
	fmt.Printf("%T\n", k)

	// Pointer of type User can also be used to call a method
	// with a value receiver.
	jill := &User{"Jill", "jill@email.com"}
	jill.Notify()
}
