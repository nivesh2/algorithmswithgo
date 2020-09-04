package main

import "fmt"

type node struct {
	prev *node
	data int
}

type list struct {
	bottom *node
}

func (l *list) push(data int) {
	temp := new(node)
	temp.data = data

	if l.bottom == nil {
		temp.prev = nil
		l.bottom = temp
	} else {
		temp.prev = l.bottom
		l.bottom = temp
	}
}

func (l *list) pop() int {
	temp := 0
	if l.bottom == nil {
		return -1
	}
	temp = l.bottom.data
	fmt.Println(temp)
	l.bottom = l.bottom.prev
	return temp
}

func main() {

	fmt.Println("main start ****************")
	l := list{}

	l.push(1)
	l.push(2)
	l.push(3)

	l.pop()
	l.pop()
	l.push(100)
	l.pop()
	fmt.Println("main complete")
}
