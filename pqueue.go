// Package implements the priority queue data structure
package pqueue

type Pqueue struct {
	Pqueue []int
}

func NewPqueue() *Pqueue {
	return &Pqueue{Pqueue: []int{0}}
}

func (p *Pqueue) percolateUp() {
	element := p.Size()

	for {

		if element <= 1 {
			break
		}

		parent := element / 2
		if p.Pqueue[parent] > p.Pqueue[element] {
			temp := p.Pqueue[parent]
			p.Pqueue[parent] = p.Pqueue[element]
			p.Pqueue[element] = temp
		} else {
			break
		}

		element = parent
	}
}

func (p *Pqueue) percolateDown() {
	element := 1
	for {
		leftChild := element * 2
		rightChild := (element * 2) + 1

		if leftChild <= p.Size() && p.Pqueue[element] > p.Pqueue[leftChild] {
			p.Pqueue[element], p.Pqueue[leftChild] = p.Pqueue[leftChild], p.Pqueue[element]
			element = leftChild
		} else if rightChild <= p.Size() && p.Pqueue[element] > p.Pqueue[rightChild] {
			p.Pqueue[element], p.Pqueue[rightChild] = p.Pqueue[rightChild], p.Pqueue[element]
			element = rightChild
		} else {
			break
		}

	}
}

func (p *Pqueue) Insert(k int) {
	p.Pqueue = append(p.Pqueue, k)
	p.percolateUp()
}

func (p Pqueue) FindMin() int {
	return p.Pqueue[1]
}

func (p *Pqueue) DelMin() int {
	min, length := p.Pqueue[1], p.Size()
	p.Pqueue[1] = p.Pqueue[length]
	p.Pqueue = p.Pqueue[:length-1]

	p.percolateDown()

	return min
}

func (p Pqueue) Size() int {
	return len(p.Pqueue) - 1
}

func (p Pqueue) IsEmpty() bool {
	return p.Size() < 1
}
