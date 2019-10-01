package main

import (
	"errors"
	"fmt"
)

type wNode struct {
	data       int
	expression func(int, int) bool
	nodeLeft   *wNode
	nodeRight  *wNode
}

// assuming true == right
// and false == left

func (w *wNode) Insert(data int) {
	newNode := wNode{
		data:       data,
		expression: w.expression,
	}

	// eval expression
	if w.expression(w.data, data) {
		if w.nodeRight != nil {
			w.nodeRight.Insert(data)
		} else {
			w.nodeRight = &newNode
		}
	} else {
		if w.nodeLeft != nil {
			w.nodeLeft.Insert(data)
		} else {
			w.nodeLeft = &newNode
		}
	}
	return
}

func (w *wNode) Remove(node *wNode) error {
	//var tempPrev *wNode

	return nil
}

func (w *wNode) Find(data int) (*wNode, error) {
	var temp *wNode
	var err error

	if w.data == data {
		return w, nil
	}

	if w.nodeRight == nil && w.nodeLeft == nil {
		err = errors.New("End of tree")
		return temp, err
	}

	if w.expression(w.data, data) {
		if w.nodeRight != nil {
			temp, err = w.nodeRight.Find(data)
		}
	} else {
		if w.nodeLeft != nil {
			temp, err = w.nodeLeft.Find(data)
		}
	}

	return temp, err
}

func main() {
	node := wNode{
		expression: func(parent int, candidate int) bool { return parent > candidate },
	}

	node.Insert(-1)
	node.Insert(12)
	node.Insert(5)
	fmt.Println(node.Find(11))
}
