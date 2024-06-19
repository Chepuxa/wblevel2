package main

import "fmt"

/*
Реализовать паттерн билдер, объяснить применимость, плюсы и минусы, а также реальные примеры использования паттерна на практике.

Builder - это паттерн, который позволяет поэтапно создавать сложные объекты с помощью четко определенной последовательности действий

Применимость:
- Когда нужно избавиться от переизбытка конструкторов для объекта с множеством полей.
- Когда код должен создавать разные представления какого-то объекта.
- Когда нужно собирать сложные составные объекты.

Плюсы:
- Позволяет создавать объекта пошагово.
- Позволяет использовать один и тот же код для создания различных объектов.
- Изолирует сложный код сборки объекта от его основной бизнес-логики.

Минусы:
- Усложняет структур проекта.
- Увеличивает шанс ошибки пользователя при работе с билдером напрямую, в случае если будут не заданы важные для логики поля.
*/

type Object struct {
	a int
	b string
	c bool
}

type ObjectBuilder struct {
	object Object
}

func NewObjectBuilder() ObjectBuilder {
	return ObjectBuilder{}
}

func (ob ObjectBuilder) setA(a int) ObjectBuilder {
	ob.object.a = a
	return ob
}

func (ob ObjectBuilder) setB(b string) ObjectBuilder {
	ob.object.b = b
	return ob
}

func (ob ObjectBuilder) setC(c bool) ObjectBuilder {
	ob.object.c = c
	return ob
}

func (ob ObjectBuilder) build() Object {
	return ob.object
}

func main() {
	object := NewObjectBuilder().
		setA(10).
		setB("object").
		setC(true).
		build()
	fmt.Println(object)
}
