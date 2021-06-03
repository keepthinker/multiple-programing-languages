package main

import "fmt"

type Person struct {
	name string
	age int
}

type Worker struct {
	person Person
	tool string
}

type Teacher struct {
	Person
	level int
}

func main() {
	var worker Person
	worker.name = "John"
	worker.age = 45
	PrintPerson(worker)

	teacher := Person{ "mask", 22}
	PrintPerson(teacher)

	engineer := Person { name: "mike", age: 49}
	PrintPerson(engineer)

	engineer.age = 50
	PrintPerson(engineer)

	maintainer := Worker{person: Person{name: "Jake", age: 30}, tool: "hammer"}
	PrintWorker(maintainer)

	teacher1 := Teacher{Person{name: "Jake", age: 30}, 1}
	PrintTeacher(teacher1)
	teacher1.Person.name = "Jackson"
	PrintTeacher(teacher1)

}


func PrintPerson(per Person) {
	fmt.Printf("name:%s, age:%d\n", per.name, per.age)
}

func PrintWorker(worker Worker) {
	fmt.Printf("name:%s, age:%d, tool:%s\n", worker.person.name, worker.person.age, worker.tool)
}

func PrintTeacher(teacher Teacher) {
	fmt.Printf("name:%s, age:%d, level:%d\n", teacher.name, teacher.age, teacher.level)
}
