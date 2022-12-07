package framework

import (
	"container/heap"
	"reflect"
)

type Heap[T Lessor[T]] []T

func (h Heap[T]) Len() int           { return len(h) }
func (h Heap[T]) Less(i, j int) bool { return h[i].Less(h[j]) }
func (h Heap[T]) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

// Push and Pop use pointer receivers because they modify the slice's length,
// not just its contents.
// TODO(ppg): allow meeting an interface to modify item on push and ppo;
// for example PriorityQueue wants to set index=len(*h) on Push and clear
// that on Pop.

func (h *Heap[T]) Push(x any) {
	item := x.(T)
	//item.index = len(*h)
	*h = append(*h, item)
}

func (h *Heap[T]) Pop() any {
	old := *h
	n := len(old)
	item := old[n-1]
	// avoid memory leak
	old[n-1] = reflect.Zero(reflect.TypeOf(item)).Interface().(T)
	//item.index = -1 // for safety
	*h = old[0 : n-1]
	return item
}

type HeapInt int

func (i HeapInt) Less(other HeapInt) bool { return i < other }

var _ heap.Interface = (*Heap[HeapInt])(nil)
