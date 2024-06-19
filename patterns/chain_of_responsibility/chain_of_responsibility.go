package main

import (
	"fmt"
)

/*
Реализовать паттерн цепочка вызовов, объяснить применимость, плюсы и минусы, а также реальные примеры использования паттерна на практике.

Цепочка вызовов — это поведенческий паттерн проектирования, который позволяет передавать запросы последовательно по цепочке обработчиков.
Каждый последующий обработчик решает, может ли он обработать запрос сам и стоит ли передавать запрос дальше по цепи.

Применимость:
- Когда программа должна обрабатывать разнообразные запросы несколькими способами, но заранее неизвестно, какие конкретно запросы будут приходить и какие обработчики для них понадобятся.
- Когда важно, чтобы обработчики выполнялись один за другим в строгом порядке.
- Когда набор объектов, способных обработать запрос, должен задаваться динамически.

Плюсы:
- Уменьшает зависимость между клиентом и обработчиками.
- Реализует принцип единственной обязанности.
- Реализует принцип открытости/закрытости.

Минусы:
- Запрос может остаться никем не обработанным.
*/

type Patient struct {
	name             string
	registrationDone bool
	doctorChekUpDone bool
	medicineDone     bool
	paymentDone      bool
}

type Departament interface {
	execute(*Patient)
	setNext(Departament)
}

type Reception struct {
	next Departament
}

func (r *Reception) execute(p *Patient) {
	if p.registrationDone {
		fmt.Println("Patient registration already done")
		r.next.execute(p)
		return
	} else {
		fmt.Println("Reception registering patient")
		p.registrationDone = true
		r.next.execute(p)
	}
}

func (r *Reception) setNext(next Departament) {
	r.next = next
}

type Doctor struct {
	next Departament
}

func (d *Doctor) execute(p *Patient) {
	if p.doctorChekUpDone {
		fmt.Println("Doctor checkup already done")
		d.next.execute(p)
		return
	} else {
		fmt.Println("Doctor checking patient")
		p.doctorChekUpDone = true
		d.next.execute(p)
	}
}

func (d *Doctor) setNext(next Departament) {
	d.next = next
}

type Medical struct {
	next Departament
}

func (m *Medical) execute(p *Patient) {
	if p.medicineDone {
		fmt.Println("Medicine already given to patient")
		m.next.execute(p)
		return
	} else {
		fmt.Println("Medical giving medicine to patient")
		p.medicineDone = true
		m.next.execute(p)
	}
}

func (m *Medical) setNext(next Departament) {
	m.next = next
}

type Cashier struct {
	next Departament
}

func (c *Cashier) execute(p *Patient) {
	if p.paymentDone {
		fmt.Println("Payment Done")
		return
	} else {
		fmt.Println("Patient paying to the cashier")
		p.paymentDone = true
	}
}

func (c *Cashier) setNext(next Departament) {
	c.next = next
}

func main() {
	cashier := &Cashier{}

	medical := &Medical{}
	medical.setNext(cashier)

	doctor := &Doctor{}
	doctor.setNext(medical)

	reception := &Reception{}
	reception.setNext(doctor)

	patient := &Patient{name: "Jacob"}
	reception.execute(patient)
}