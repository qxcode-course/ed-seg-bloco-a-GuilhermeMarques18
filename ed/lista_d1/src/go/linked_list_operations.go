func (l *LList) PushBack(value int) {
	newNode := &Node{
		Value: value,
		prev:  l.root.prev,
		next:  l.root,
		root:  l.root,
	}

	l.root.prev.next = newNode
	l.root.prev = newNode
	l.size++
}

func (l *LList) PopBack() {
	if l.size == 0 {
		return
	}

	last := l.root.prev

	last.prev.next = l.root
	l.root.prev = last.prev

	l.size--
}
