package main

import (
	"fmt"
)

/*
Реализовать паттерн посетитель, объяснить применимость, плюсы и минусы, а также реальные примеры использования паттерна на практике.

Посетитель — это поведенческий паттерн проектирования, который позволяет добавлять в программу новые операции, не изменяя классы объектов, над которыми эти операции могут выполняться.

Применимость:
- Когда нужно выполнить одну и ту же операцию для ряда объектов, без определения этой операции в объекте.
- Когда нужно определить новую операцию над объектами без изменения их структуры.

Плюсы:
- Упрощает добавление операций, работающих со сложными структурами объектов.
- Объединяет родственные операции в одном классе.
- Посетитель может накапливать состояние при обходе структуры элементов.

Минусы:
- Паттерн не оправдан, если иерархия элементов часто меняется.
- Может привести к нарушению инкапсуляции элементов.
*/

type ObjectA struct {
	a int
	b int
}

func (objA *ObjectA) accept(v Visitor) {
	v.visitObjectA(objA)
}

type ObjectB struct {
	a int
	b int
}

func (objB *ObjectB) accept(v Visitor) {
	v.visitObjectB(objB)
}

type Visitor interface {
	visitObjectA(*ObjectA)
	visitObjectB(*ObjectB)
}

type ObjectCalculator struct {
}

func (ac *ObjectCalculator) visitObjectA(objA *ObjectA) {
	result := objA.a + objA.b
	fmt.Printf("Result: %v\b", result)
}

func (ac *ObjectCalculator) visitObjectB(objB *ObjectB) {
	result := objB.a * 2 + objB.b * 2
	fmt.Printf("Result: %v\b", result)
}

func main() {
	objA := &ObjectA{3, 4}
	objB := &ObjectB{3, 4}
	objCalc := &ObjectCalculator{}
	objA.accept(objCalc)
	objB.accept(objCalc)
}