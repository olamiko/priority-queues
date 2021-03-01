// Package implements the priority queue data structure
package pqueue

type BinaryHeap func() []int

type Pqueue struct {
	Pqueue BinaryHeap
}

func BinaryHeap() []int{
	return []int{0}
}

func (p *Pqueue) percolateUp() {
	element := p.Size()

	for {

		if element <= 1 {
			break
		}

		parent := element / 2
		if p.BinaryHeap[parent] > p.BinaryHeap[element] {
			temp := p.BinaryHeap[parent]
			p.BinaryHeap[parent] = p.BinaryHeap[element]
			p.BinaryHeap[element] = temp
		} else {
			break
		}

		element = parent
	}
}

func (p *Pqueue) percolateDown() {
	element = 1
	for {
		leftChild = element * 2

		if p.Size() <= leftChild {
			break
		}
		if p.Size() <= rightChild {
			rightChild = element
		}

		if p.BinaryHeap[element] > p.BinaryHeap[leftChild] {
			p.BinaryHeap[element], p.BinaryHeap[leftChild] = p.BinaryHeap[leftChild], p.BinaryHeap[element]
			element = leftChild
		} else if p.BinaryHeap[element] > p.BinaryHeap[rightChild] {
			p.BinaryHeap[element], p.BinaryHeap[rightChild] = p.BinaryHeap[rightChild], p.BinaryHeap[element]
			element = rightChild
		} else {
			break
		}

	}
}

func (p *Pqueue) Insert(k) {
	p.BinaryHeap = append(p.BinaryHeap, k)
	p.percolateUp()
}

func (p Pqueue) FindMin() int {
	return p.BinaryHeap[0]
}

func (p *Pqueue) DelMin() int {
	min := p.BinaryHeap[1]
	length := p.Size()
	p.BinaryHeap[0] = p.BinaryHeap[length]
	p.BinaryHeap = p.BinaryHeap[:-1]

	p.percolateDown()

	return min
}

func (p Pqueue) Size() int {
	return len(p.BinaryHeap) - 1
}

func (p pqueue) IsEmpty() bool {
	return p.size() > 0
}
