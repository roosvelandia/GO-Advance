package main

import "fmt"

// patron de comportamientos,

// create an interface to register the observers and the broadcast method
type Topic interface {
	register(observer Observer)
	broadcast()
}

// create my object observer
type Observer interface {
	getId() string
	updateValue(string)
}

// item -< No disponible
// item -Z avise -> hay disponible
// each item will have an split of observers
type Item struct {
	observers []Observer
	name      string
	available bool
}

//create an object of type item
func NewItem(name string) *Item {
	return &Item{
		name: name,
	}
}

//update availability
func (i *Item) UpdateAvailable() {
	fmt.Printf("Item %s is available  \n", i.name)
	i.available = true
	i.broadcast()
}

// register observers
func (i *Item) register(observer Observer) {
	i.observers = append(i.observers, observer)
}

// for each observer, send a notification
func (i *Item) broadcast() {
	for _, observer := range i.observers {
		observer.updateValue(i.name)
	}
}

// type of observer
type EmailClient struct {
	id string
}

// update value for observer
func (eC *EmailClient) updateValue(value string) {
	fmt.Printf("Sending email - %s available from client %s  \n", value, eC.id)
}
func (eC *EmailClient) getId() string {
	return eC.id
}

func main() {
	nvidiaItem := NewItem("RTX 3080")
	firstObserver := &EmailClient{
		id: "12AB",
	}
	secondObserver := &EmailClient{
		id: "34CD",
	}
	nvidiaItem.register(firstObserver)
	nvidiaItem.register(secondObserver)
	// when i do this the item is now available so observers will receive the email
	nvidiaItem.UpdateAvailable()
}
