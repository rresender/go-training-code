package main

import (
	"fmt"
	"math"
	"sort"
)

type IceCream struct {
	index  int
	flavor int
}

type IceCreams []IceCream

func (s IceCreams) Len() int {
	return len(s)
}

func (s IceCreams) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s IceCreams) Less(i, j int) bool {
	return s[i].index < s[j].index
}

func binarySearch(values []IceCream, init, end, search int) int {

	m := (end + init) / 2

	if values[m].index == search {
		return m
	}

	if init == end {
		return -1
	}

	if values[m].index < search {
		return binarySearch(values, init+1, end, search)
	}

	return binarySearch(values, init, end-1, search)
}

func main() {

	var t int
	fmt.Scanf("%d", &t)

	for trip := 0; trip < t; trip++ {
		var m int
		fmt.Scanf("%d", &m)

		var n int
		fmt.Scanf("%d", &n)

		values := make(IceCreams, n)

		for i := 0; i < n; i++ {
			var v int
			fmt.Scanf("%d", &v)
			values[i] = IceCream{v, i + 1}
		}

		values = sort.Interface(values).(IceCreams)

		for i := 0; i < n-1; i++ {
			search := m - values[i].flavor
			if search >= values[i].flavor {
				index := binarySearch(values, i+1, n-1, search)
				if index != -1 {
					min := math.Min(float64(values[i].index), float64(index))
					max := math.Max(float64(values[i].index), float64(index))
					fmt.Printf("%f %f\n", min, max)
					break

				}
			}
		}

	}

}
