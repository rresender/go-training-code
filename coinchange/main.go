package main

import (
	"fmt"
)

func main() {

	var n int
	fmt.Scanf("%d", &n)

	var m int
	fmt.Scanf("%d", &m)

	coins := make([]int, m)

	for m > 0 {
		var coin int
		fmt.Scanf("%d", &coin)
		coins = append(coins, coin)
		m--
	}

	combinations := make([]int, n+1)

	combinations[0] = 1

	for _, coin := range coins {
		for i := 1; i < len(combinations); i++ {
			if i >= coin {
				combinations[i] += combinations[i-coin]
			}
		}
	}

	fmt.Println(combinations[n])

}
