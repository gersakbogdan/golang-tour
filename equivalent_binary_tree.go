package main

import (
	"fmt"
	"math/rand"
)

type Tree struct {
	Left *Tree
	Value int
	Right *Tree
}

func New(n, k int) *Tree {
	var t *Tree

	for _, v := range rand.Perm(n) {
		t = insert(t, (1 + v) * k)
	}

	return t
}

func insert(t *Tree, v int) *Tree {
	if t == nil {
		return &Tree{nil, v, nil}
	}
	if v < t.Value {
		t.Left = insert(t.Left, v)
		return t
	}
	t.Right = insert(t.Right, v)
	return t;
}

func Walk(t *Tree, c chan int) {
	if t == nil {
		return
	}

	Walk(t.Left, c)
	c <- t.Value
	Walk(t.Right, c)
}

func Walker(t *Tree) <-chan int {
	c := make(chan int)
	go func() {
		Walk(t, c)
		close(c)
	}()
	return c
}

func Same(t1 *Tree, t2 *Tree) bool {
	ch1, ch2 := Walker(t1), Walker(t2)

	for {
		v1, ok1 := <-ch1
		v2, ok2 := <-ch2

		if !ok1 || !ok2 {
			return ok1 == ok2
		}

		if v1 != v2 {
			break
		}
	}

	return false
}

func main() {
	fmt.Println(Same(New(10, 1), New(10, 2)))
}