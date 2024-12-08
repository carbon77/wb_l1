package main

import "fmt"

// Интерфейс USB для взаимодействия с компьютером
type USB interface {
	connectWithUsbCable()
}

// Класс карты памяти. Он не обладает интерфейсом для взаимодействия с компьютером,
// поэтому его необходимо адаптировать с помощью картридера. Является адаптируемым классов (Adaptee)
type MemoryCard struct {
}

func (mc *MemoryCard) insert() {
	fmt.Println("Карта памяти вставлена")
}

func (mc *MemoryCard) copyData() {
	fmt.Println("Данные скопированы")
}

// Класс картридера, является адаптером (Adapter). Реализует интерфейс USB для взаимодействия с компьютером
// С помощью композиции использует класс MemoryCard для его адаптации
type CardReader struct {
	memoryCard *MemoryCard
}

func (cr *CardReader) connectWithUsbCable() {
	cr.memoryCard.insert()
	cr.memoryCard.copyData()
}

func main() {
	var cardReader USB
	cardReader = &CardReader{
		memoryCard: &MemoryCard{},
	}

	cardReader.connectWithUsbCable()
}
