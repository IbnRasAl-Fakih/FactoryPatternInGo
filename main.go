package main

import "fmt"

type IPhone interface {
	getName() string
	getBatteryCapacity() int
	getMemoryCapacity() int
}

type Phone struct {
	name    string
	battery int
	memory  int
}

func (r *Phone) getName() string {
	return r.name
}

func (r *Phone) getBatteryCapacity() int {
	return r.battery
}

func (r *Phone) getMemoryCapacity() int {
	return r.memory
}

type Redmi struct {
	Phone
}

func newRedmi() IPhone {
	return &Redmi{
		Phone: Phone{
			name:    "Redmi Note 12 Pro",
			battery: 5000,
			memory:  256,
		},
	}
}

type Oppo struct {
	Phone
}

func newOppo() IPhone {
	return &Oppo{
		Phone: Phone{
			name:    "Oppo A96",
			battery: 4800,
			memory:  128,
		},
	}
}

func getPhone(Phone string) IPhone {
	if Phone == "Oppo" {
		return newOppo()
	}
	if Phone == "Redmi" {
		return newRedmi()
	}

	return nil
}

func main() {
	Redmi := getPhone("Redmi")
	Oppo := getPhone("Oppo")

	fmt.Println(Redmi.getName()+" - battery capacity: ", Redmi.getBatteryCapacity(), " , memory capacity: ", Redmi.getMemoryCapacity())
	fmt.Println(Oppo.getName()+" - battery capacity: ", Oppo.getBatteryCapacity(), " , memory capacity: ", Oppo.getMemoryCapacity())
}
