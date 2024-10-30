package main

import "fmt"

type livingBeing interface {
	isLivingCreature() bool
}
type dog struct {
	breed   string
	isAlive bool
}

type human struct {
	race    string
	isAlive bool
}

func (d dog) isLivingCreature() bool {
	return d.isAlive
}

func (h human) isLivingCreature() bool {
	return h.isAlive
}

func main() {
	jozo := human{"white", true}
	jazvecik := dog{"jazvecik", true}

	livingBeings := []livingBeing{jozo, jazvecik}
	fmt.Println(livingBeings[1].isLivingCreature())

}
