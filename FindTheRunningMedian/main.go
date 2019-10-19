package main

import (
	"container/heap"
	"fmt"
	"strconv"
)

type MinHeap []int

func (h MinHeap) Len() int            { return len(h) }
func (h MinHeap) Less(i, j int) bool  { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(int)) }

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h MinHeap) Peek() interface{} {
	return h[0]
}

type MaxHeap struct {
	MinHeap
}

func (h MaxHeap) Less(i, j int) bool { return h.MinHeap[i] > h.MinHeap[j] }

func main() {
	var str string
	fmt.Scanln(&str)
	n, _ := strconv.Atoi(str)

	l := &MinHeap{}
	heap.Init(l)

	h := &MaxHeap{}
	heap.Init(h)

	for n > 0 {
		fmt.Scanln(&str)
		v, _ := strconv.Atoi(str)

		if l.Len() == 0 || l.Peek().(int) < v {
			heap.Push(l, v)
		} else {
			heap.Push(h, v)
		}
		if l.Len()-h.Len() >= 2 {
			heap.Push(h, heap.Pop(l))
		} else if h.Len()-l.Len() >= 2 {
			heap.Push(l, heap.Pop(h))
		}

		if l.Len() == h.Len() {
			sum := float64(l.Peek().(int) + h.Peek().(int))
			div := sum / 2
			fmt.Printf("%.1f\n", div)
		} else {
			if l.Len() > h.Len() {
				v := l.Peek().(int)
				fmt.Printf("%.1f\n", float64(v))
			} else {
				v := h.Peek().(int)
				fmt.Printf("%.1f\n", float64(v))
			}
		}
		n--
	}
}
