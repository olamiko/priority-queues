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

	for (element / 2) > 0{

		if p.Pqueue[element/2] > p.Pqueue[element] {
			p.Pqueue[element], p.Pqueue[element/2] = p.Pqueue[element/2], p.Pqueue[element]
		} else {
			break
		}

		element = element/2
	}
}

func (p *Pqueue) minChild(index int) int{
	if (index * 2) + 1 > p.Size(){
		return index * 2
	}

	if p.Pqueue[index * 2] < p.Pqueue[(index * 2) + 1]{
		return index * 2
	}

	return (index * 2) + 1
}

func (p *Pqueue) percolateDown() {
	element := 1
	for element * 2 <= p.Size() {
		minChild := p.minChild(element)

		if p.Pqueue[element] > p.Pqueue[minChild] {
			p.Pqueue[element], p.Pqueue[minChild] = p.Pqueue[minChild], p.Pqueue[element]
			element = minChild
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
