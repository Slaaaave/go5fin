package personaldata

import "fmt"

type Personal struct {
	Name   string
	Weight float64
	Height int
	// TODO: добавить поля
}

func (p Personal) Print() {
	// TODO: реализовать функцию
	fmt.Println("Имя:", p.Name)
	fmt.Println("Вес:", p.Weight)
	fmt.Println("Рост:", p.Height)
}
