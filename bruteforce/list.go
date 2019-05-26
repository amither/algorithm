package bruteforce

import "fmt"

type listnode struct {
	v    int
	next *listnode
}

type list struct {
	head *listnode
	tail *listnode
}

func (l *list) insert(i int) {
	a := listnode{v: i, next: nil}
	if l.head == nil {
		l.tail = &a
		l.head = l.tail
	} else {
		l.tail.next = &a
		l.tail = &a
	}
}

func (l *list) travel() {
	p := l.head
	for ; p != nil; p = p.next {
		fmt.Println(p.v)
	}
}

func (l *list) reverse() {
	var n, p, q *listnode
	n = nil
	p = l.head
	q = l.head
	l.tail = l.head
	for p != nil {
		q = p.next
		p.next = n
		n = p
		p = q
	}
	l.head = n

}
