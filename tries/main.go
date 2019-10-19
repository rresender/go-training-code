package main

import (
	"fmt"
	"strconv"
)

const alphabetSize = 26

type Node struct {
	children    []*Node
	isEndOfWord bool
	size        int
}

var root *Node

func getNode() *Node {
	n := &Node{}
	n.children = make([]*Node, alphabetSize)
	return n
}

func add(s string) {
	crawl := root
	for _, c := range s {
		index := c - 'a'
		if crawl.children[index] == nil {
			crawl.children[index] = getNode()
		}
		crawl = crawl.children[index]
		crawl.size++
	}
	crawl.isEndOfWord = true
}

func find(s string) int {
	crawl := root
	for _, c := range s {
		index := c - 'a'
		if crawl.children[index] == nil {
			return 0
		}
		crawl = crawl.children[index]
	}
	return crawl.size
}

func main() {

	var in string
	fmt.Scanln(&in)
	n, _ := strconv.Atoi(in)

	root = getNode()

	for n > 0 {

		var cmd string
		fmt.Scanf("%s", &cmd)

		var value string
		fmt.Scanf("%s", &value)

		switch cmd {
		case "add":
			add(value)
		case "find":
			fmt.Printf("%d\n", find(value))
		}

		n--
	}

}
