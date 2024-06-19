package main

import (
	"fmt"
)

/*
Реализовать паттерн стратегия, объяснить применимость, плюсы и минусы, а также реальные примеры использования паттерна на практике.

Стратегия — это поведенческий паттерн проектирования, который определяет семейство схожих алгоритмов и помещает каждый из них в собственный класс,
после чего алгоритмы можно взаимозаменять прямо во время исполнения программы.

Применимость:
- Когда нужно использовать разные вариации какого-то алгоритма внутри одного объекта.
- Когда есть множество похожих классов, отличающихся только некоторым поведением.
- Когда нужно не обнажать детали реализации алгоритмов для других классов.
- Когда различные вариации алгоритмов реализованы в виде развесистого условного оператора. Каждая ветка такого оператора представляет собой вариацию алгоритма.

Плюсы:
- Горячая замена алгоритмов на лету.
- Изолирует код и данные алгоритмов от остальных классов.
- Уход от наследования к делегированию.
- Реализует принцип открытости/закрытости.

Минусы:
- Усложняет программу за счёт дополнительных классов.
- Клиент должен знать, в чём состоит разница между стратегиями, чтобы выбрать подходящую.
*/

type PaymentStrategy interface {
	Pay(amount float64) string
}

type CreditCardPayment struct {
}

func (p *CreditCardPayment) Pay(amount float64) string {
	return fmt.Sprintf("Paid %.2f $ using credit card", amount)
}

type PayPalPayment struct {
}

func (p *PayPalPayment) Pay(amount float64) string {
	return fmt.Sprintf("Paid %.2f $ using PayPal", amount)
}

type PaymentContext struct {
	strategy PaymentStrategy
}

func (c *PaymentContext) SetStrategy(strategy PaymentStrategy) {
	c.strategy = strategy
}

func (c *PaymentContext) MakePayment(amount float64) string {
	return c.strategy.Pay(amount)
}

func main() {
	creditCardStrategy := CreditCardPayment{}
	payPalStrategy := PayPalPayment{}

	context := PaymentContext{}
	context.SetStrategy(&creditCardStrategy)
	fmt.Println(context.MakePayment(100.48))

	context.SetStrategy(&payPalStrategy)
	fmt.Println(context.MakePayment(44.88))
}