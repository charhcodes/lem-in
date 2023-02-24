package main

type Node struct {
	name string // holds a string
	next *Node  // points to next node in the List
}

type List struct {
	head *Node // points to first node in the List
}

// adds a new node with the given data to the List
func (l *List) Add(name string) {
	newNode := &Node{name: name}
	// creates a new Node object and assigns it to newNode
	// this is so we can add newNode to a linked list

	if l.head == nil {
		l.head = newNode // if List is empty, new node becomes the first node
	} else {
		current := l.head
		for current.next != nil { // if current node is last
			current = current.next // current node = next node
		}
		current.next = newNode // next node = last node
	}
}

// removes node with specified data from the list
func (l *List) Remove(name string) {
	if l.head == nil { // if list is empty, return
		return
	}

	if l.head.name == name { // if head == data, head points to next node
		l.head = l.head.next
		return
	}

	current := l.head // current == first node in list
	for current.next != nil {
		if current.next.name == name { // if node after current == data
			current.next = current.next.next // updates 'next' field of current node to point to node after it
			// effectively deletes the next node from the list eg. '1 2 3 4' -> '1 3 4' if 1 = current
			return
		}
		current = current.next // current = next node so loop can go again
	}
}

// checks if a node with the given data is in the list
func (l *List) Contains(name string) bool {
	current := l.head
	for current != nil { // when current is not empty
		if current.name == name { // if current == data
			return true // return true means list already contains the current node
		}
		current = current.next // current = next node, loop starts again
	}
	return false
}
