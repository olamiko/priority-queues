package pqueue

import "testing"

var pqueue = NewPqueue()

func TestSize(t *testing.T){
	description := "Test priority queue size function"

	if pqueue.Size() != 0 {
		t.Fatalf("FAIL: %s\nExpected: %#v\nActual: %#v", description, 0, pqueue.Size())
	}

}

func TestInsert(t *testing.T){

	description := "Test priority queue insert function"
	
	pqueue.Insert(5)
	pqueue.Insert(10)
	pqueue.Insert(2)

	var actual [4]int
	copy(actual[:], pqueue.Pqueue[0:4])
	expected := [...]int{0,2,10,5}

	if actual != expected  {
		t.Fatalf("FAIL: %s\nExpected: %#v\nActual: %#v", description, expected, actual)
	}
}

func TestFindMin(t *testing.T){

	description := "Test priority queue findmin function"
	
	if pqueue.FindMin() != 2  {
		t.Fatalf("FAIL: %s\nExpected: %#v\nActual: %#v", description, 2, pqueue.FindMin())
	}
}


func TestDelMin(t *testing.T){

	description := "Test priority queue delmin function"

	pqueue.Insert(1)

	min := pqueue.DelMin() 

	if min != 1 {
		t.Fatalf("FAIL: %s\nExpected: %#v\nActual: %#v", description, 2, min)
	}
	
	var actual [4]int
	copy(actual[:], pqueue.Pqueue[0:4])
	expected := [...]int{0,2,10,5}

	if actual != expected  {
		t.Fatalf("FAIL: %s\nExpected: %#v\nActual: %#v", description, expected, actual)
	}

}

func TestIsEmpty(t *testing.T){

	description := "Test priority queue isempty function"
	
	if pqueue.IsEmpty() {
		t.Fatalf("FAIL: %s\nExpected: %#v\nActual: %#v", description, false, pqueue.IsEmpty())
	}
}


