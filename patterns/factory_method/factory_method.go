package main

import (
	"fmt"
)

/*
Реализовать паттерн фабричнрый метод, объяснить применимость, плюсы и минусы, а также реальные примеры использования паттерна на практике.

Фабричный метод — это порождающий паттерн проектирования, который определяет общий интерфейс для создания объектов в суперклассе,
позволяя подклассам изменять тип создаваемых объектов.

Применимость:
- Когда заранее неизвестны типы и зависимости объектов, с которыми должен работать код.
- Когда нужно дать возможность пользователям расширять части фреймворка или библиотеки.
- Когда нужно экономить системные ресурсы, повторно используя уже созданные объекты, вместо порождения новых.

Плюсы:
- Избавляет класс от привязки к конкретным классам продуктов.
- Выделяет код производства продуктов в одно место, упрощая поддержку кода.
- Упрощает добавление новых продуктов в программу.
- Реализует принцип открытости/закрытости.

Минусы:
- Может привести к созданию больших параллельных иерархий классов, так как для каждого класса продукта надо создать свой подкласс создателя.
*/

type IGun interface {
	setName(name string)
	setPower(power int)
	getName() string
	getPower() int
}

type Gun struct {
	name  string
	power int
}

func (g *Gun) setName(name string) {
	g.name = name
}

func (g *Gun) setPower(power int) {
	g.power = power
}

func (g *Gun) getName() string {
	return g.name
}

func (g *Gun) getPower() int {
	return g.power
}

type BFG struct {
	Gun
}

func newBFG() IGun {
	return &BFG{
		Gun: Gun{
			name:  "Big Fucking Gun",
			power: 9,
		},
	}
}

type Buriza struct {
	Gun
}

func newBuriza() IGun {
	return &Buriza{
		Gun: Gun{
			name:  "Buriza-Do Kyanon",
			power: 5,
		},
	}
}

func getGun(gunType string) (IGun, error) {
	if gunType == "BFG" {
		return newBFG(), nil
	} else if gunType == "Buriza" {
		return newBuriza(), nil
	} else {
		return nil, fmt.Errorf("%v - is wrong gun type", gunType)
	}
}

func printDetails(g IGun) {
	fmt.Printf("Gun's name: %v\n", g.getName())
	fmt.Printf("Gun's power: %v\n", g.getPower())
}

func main() {
	bfg, err := getGun("BFG")
	if err != nil {
		fmt.Println(err)
	} else {
		printDetails(bfg)
	}

	buriza, err := getGun("Buriza")
	if err != nil {
		fmt.Println(err)
	} else {
		printDetails(buriza)
	}

	rocketLauncher, err := getGun("Rocket")
	if err != nil {
		fmt.Println(err)
	} else {
		printDetails(rocketLauncher)
	}
}