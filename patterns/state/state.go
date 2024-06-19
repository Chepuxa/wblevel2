package main

import (
	"fmt"
)

/*
Реализовать паттерн состояние, объяснить применимость, плюсы и минусы, а также реальные примеры использования паттерна на практике.

Состояние — это поведенческий паттерн проектирования, который позволяет объектам менять поведение в зависимости от своего состояния.
Извне создаётся впечатление, что изменился класс объекта.

Применимость:
- Когда есть объект, поведение которого кардинально меняется в зависимости от внутреннего состояния, причём типов состояний много, и их код часто меняется.
- Когда код класса содержит множество больших, похожих друг на друга, условных операторов, которые выбирают поведения в зависимости от текущих значений полей класса.
- Когда используется табличная машина состояния, построенная на условных операторах, и возникает дублирование кода для похожих состояний и переходов.

Плюсы:
- Избавляет от множества больших условных операторов машины состояний.
- Концентрирует в одном месте код, связанный с определённым состоянием.
- Упрощает код контекста.

Минусы:
- Может неоправданно усложнить код, если состояний мало и они редко меняются.
*/

type DoorState interface {
	Open()
	Close()
	Lock()
	Unlock()
}

type OpenState struct{}

func (s *OpenState) Open() {
	fmt.Println("The door is already open")
}

func (s *OpenState) Close() {
	fmt.Println("Closing the door")
}

func (s *OpenState) Lock() {
	fmt.Println("Can't lock an opened door")
}

func (s *OpenState) Unlock() {
	fmt.Println("The door is already unlocked")
}

type CloseLockState struct{}

func (s *CloseLockState) Open() {
	fmt.Println("Cant' open an locked door")
}

func (s *CloseLockState) Close() {
	fmt.Println("The door is already closed")
}

func (s *CloseLockState) Lock() {
	fmt.Println("The door is already locked")
}

func (s *CloseLockState) Unlock() {
	fmt.Println("Unlocking the door")
}

type CloseUnlockState struct{}

func (s *CloseUnlockState) Open() {
	fmt.Println("Opening the door")
}

func (s *CloseUnlockState) Close() {
	fmt.Println("The door is already closed")
}

func (s *CloseUnlockState) Lock() {
	fmt.Println("Locking the door")
}

func (s *CloseUnlockState) Unlock() {
	fmt.Println("The door is already unlocked")
}

type Door struct {
	state DoorState
}

func (d *Door) setState(state DoorState) {
	d.state = state
}

func (d *Door) Open() {
	d.state.Open()
	if _, isCloseUnlock := d.state.(*CloseUnlockState); isCloseUnlock {
		d.setState(&OpenState{})
	}
}

func (d *Door) Close() {
	d.state.Close()
	if _, isOpen := d.state.(*OpenState); isOpen {
		d.setState(&CloseUnlockState{})
	}
}

func (d *Door) Lock() {
	d.state.Lock()
	if _, isCloseUnlock := d.state.(*CloseUnlockState); isCloseUnlock {
		d.setState(&CloseLockState{})
	}
}

func (d *Door) Unlock() {
	d.state.Unlock()
	if _, isCloseLock := d.state.(*CloseLockState); isCloseLock {
		d.setState(&CloseUnlockState{})
	}
}

func main() {
	door := &Door{state: &OpenState{}}

	// Открытое состояние двери
	door.Open()
	door.Lock()
	door.Unlock()
	door.Close()

	// Вызов door.Close() изменил состояние двери на "закрыта, незаблокирована"
	door.Close()
	door.Unlock()
	door.Lock()

	// Вызов door.Lock() изменил состояние двери на "закрыта, заблокирована"
	door.Open()
	door.Close()
	door.Lock()
	door.Unlock()

	// Вызов door.Unlock() изменил состояние двери на "закрыта, незаблокирована"
	door.Open()
	// Вызов door.Open() изменил состояние двери на "открытая"
}