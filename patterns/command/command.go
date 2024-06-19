package main

import (
	"fmt"
)

/*
Реализовать паттерн команда, объяснить применимость, плюсы и минусы, а также реальные примеры использования паттерна на практике.

Команда — это поведенческий паттерн проектирования, который превращает запросы в объекты,
позволяя передавать их как аргументы при вызове методов, ставить запросы в очередь, логировать их, а также поддерживать отмену операций.

Применимость:
- Когда нужно параметризовать объекты выполняемым действием.
- Когда нужно ставить операции в очередь, выполнять их по расписанию или передавать по сети.
- Когда нужна операция отмены.

Плюсы:
- Убирает прямую зависимость между объектами, вызывающими операции, и объектами, которые их непосредственно выполняют.
- Позволяет реализовать простую отмену и повтор операций.
- Позволяет реализовать отложенный запуск операций.
- Позволяет собирать сложные команды из простых.
- Реализует принцип открытости/закрытости.

Минусы:
- Усложняет код программы из-за введения множества дополнительных классов.
*/

type Command interface {
	Execute()
}

type Button struct {
	command Command
}

func (b *Button) press() {
	b.command.Execute()
}

type Device interface {
	On()
	Off()
	Next()
}

type OnCommand struct {
	device Device
}

func (c *OnCommand) Execute() {
	c.device.On()
}

type OffCommand struct {
	device Device
}

func (c *OffCommand) Execute() {
	c.device.Off()
}

type NextCommand struct {
	device Device
}

func (c *NextCommand) Execute() {
	c.device.Next()
}

type TV struct {
	isRunning bool
}

func (t *TV) On() {
	t.isRunning = true
	fmt.Println("Turning tv on")
}

func (t *TV) Off() {
	t.isRunning = false
	fmt.Println("Turning tv off")
}

func (t *TV) Next() {
	if t.isRunning {
		fmt.Println("Switching tv channel")
	} else {
		fmt.Println("Can't switch tv channel, tv off")
	}
}

func main() {
	tv := &TV{}

	onCommand := &OnCommand{device: tv}
	offCommand := &OffCommand{device: tv}
	nextCommand := &NextCommand{device: tv}

	onButton := &Button{command: onCommand}
	offButton := &Button{command: offCommand}
	nextButton := &Button{command: nextCommand}

	onButton.press()
	nextButton.press()
	offButton.press()
	nextButton.press()
}