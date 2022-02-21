package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	value int
	max   int
	next  *Node
}

var size = 0
var stack = new(Node)

func Push(v int) bool {
	if stack == nil {
		stack = &Node{v, v, nil}
		size = 1
		return true
	}

	if v > stack.max {
		temp := &Node{v, v, nil}
		temp.next = stack
		stack = temp
		size++
		return true
	} else {
		temp := &Node{v, stack.max, nil}
		temp.next = stack
		stack = temp
		size++
		return true
	}
}

func Pop(t *Node) (int, bool) {
	if size == 0 {
		return 0, false
	}

	if size == 1 {
		size = 0
		stack = nil
		return t.value, true
	}

	stack = stack.next
	size--
	return t.value, true
}

func main() {
	stack = nil

	f, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		switch {
		case strings.Contains(scanner.Text(), "push"):
			s := strings.Split(scanner.Text(), " ")
			d, _ := strconv.Atoi(s[1])
			Push(d)
		case strings.Contains(scanner.Text(), "pop"):
			_, _ = Pop(stack)
		case strings.Contains(scanner.Text(), "max"):
			fmt.Println(stack.max)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
