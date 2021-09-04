package traffic

import "fmt"

type Car struct {
	Name string
	Type int8
}

func (c *Car) DoAction() {
	fmt.Printf("I am running|name:%s|type:%d\n", c.Name, c.Type)
}

func (c *Car) ChangeName(newName string) {
	c.Name = newName
}

func CarStructOperation() {
	car := Car{
		Name: "benz",
		Type: 1,
	}
	car.ChangeName("Porsche")
	car.DoAction()

	var car1 = car
	var car2 = &car
	fmt.Printf("car: %p, car1: %p, car2: %p\n", &car, &car1, car2)
	car1.ChangeName("BMW")
	car.DoAction()
	car1.DoAction()
	car2.DoAction()

	var arr = []int{1, 2}
	for _, v := range arr {
		fmt.Printf("v: %d\n", v)
	}

	car3 := Car{Name: "Porsche", Type: 1}
	fmt.Printf("compare car struct, result:%v\n", car == car3)

}
