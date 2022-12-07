package framework

type Lessor[T any] interface {
	Less(T) bool
}

type Array[T Lessor[T]] []T

func (a Array[T]) Len() int           { return len(a) }
func (a Array[T]) Less(i, j int) bool { return a[i].Less(a[j]) }
func (a Array[T]) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
