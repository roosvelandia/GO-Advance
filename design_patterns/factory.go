package main

import "fmt"

// main definition for behaviour of an object
type IProduct interface {
	setStock(stock int)
	getStock() int
	setName(name string)
	getName() string
}

// main properties for an object
type Computer struct {
	name  string
	stock int
}

// methods implementation
func (c *Computer) setStock(stock int) {
	c.stock = stock
}
func (c *Computer) setName(name string) {
	c.name = name
}
func (c *Computer) getName() string {
	return c.name
}

func (c *Computer) getStock() int {
	return c.stock
}

// specific type of object
type Laptop struct {
	Computer
}

// factory of an objectv laptop
func newLaptop() IProduct {
	return &Laptop{
		Computer: Computer{
			name:  "Laptop Computer",
			stock: 25,
		},
	}
}

type Desktop struct {
	Computer
}

// factory of an objectv Desktop
func newDesktop() IProduct {
	return &Desktop{
		Computer: Computer{
			name:  "Desktop Computer",
			stock: 35,
		},
	}
}

//Factory, handles the construction of an object according to an input
func GetComputerFactory(computerType string) (IProduct, error) {
	if computerType == "laptop" {
		return newLaptop(), nil
	}

	if computerType == "desktop" {
		return newDesktop(), nil
	}

	return nil, fmt.Errorf("Invalid computer type")
}

func printNameAndStock(p IProduct) {
	fmt.Printf("you have a product : %s, with and stock of: %d  \n", p.getName(), p.getStock())
}

func main() {
	// factory of an object type laptop
	laptop, _ := GetComputerFactory("laptop")
	printNameAndStock(laptop)
	// factory of an object type desktop
	desktop, _ := GetComputerFactory("desktop")
	printNameAndStock(desktop)
	// factory error
	_, err := GetComputerFactory("error")
	fmt.Println(err)
}
