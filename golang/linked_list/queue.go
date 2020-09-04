package main

import (
	"errors"
	"fmt"
)

type node struct {
	next *node
	data int
}

type list struct {
	head *node
	tail *node
}

func (l *list) enqueue(data int) {
	temp := new(node)
	temp.data = data

	if l.tail == nil {
		l.tail = temp
		l.head = temp
	}
	l.tail.next = temp
	l.tail = temp
}

func (l *list) dequeue() (int, error) {
	if l.head == nil {
		return 0, errors.New("queue is empty")
	}
	temp := l.head
	l.head = l.head.next
	return temp.data, nil
}

func main() {
	fmt.Println("*********** Queue ************")
	l := new(list)
	l.enqueue(1)
	l.enqueue(2)
	l.enqueue(3)
	l.enqueue(4)
	l.enqueue(5)

	fmt.Println(l.dequeue())
	fmt.Println(l.dequeue())
	fmt.Println(l.dequeue())
	fmt.Println(l.dequeue())
	fmt.Println(l.dequeue())
	fmt.Println("compelete")
}
