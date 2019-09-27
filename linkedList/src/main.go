package main

import (
	"errors"
	"fmt"
)

type wNode struct {
	data int
	prev *wNode
	next *wNode
}

func (w *wNode) PushHead(data int) *wNode {
	newNode := &wNode{
		data: data,
	}

	w.prev = newNode

	newNode.next = w
	return newNode
}

func (w *wNode) PushTail(data int) {
	newNode := &wNode{
		data: data,
	}

	if w.next == nil {
		w.next = newNode
	} else {
		var tempPrev *wNode = nil
		tempNext := w.next

		for tempNext.next != nil {
			tempPrev = tempNext
			tempNext = tempNext.next
		}

		tempNext.prev = tempPrev
		tempNext.next = newNode
	}
}

func (w *wNode) Find(data int) (*wNode, error) {

	if w.data == data {
		return w, nil
	}

	if w.prev != nil {
		tempPrev := w.prev
		for tempPrev != nil {
			if tempPrev.data == data {
				return tempPrev, nil
			}

			tempPrev = tempPrev.prev
		}
	}

	if w.next != nil {
		tempNext := w.next
		for tempNext != nil {
			if tempNext.data == data {
				return tempNext, nil
			}

			tempNext = tempNext.next
		}
	}

	return nil, errors.New("Data does not exist in list")
}

func (w *wNode) Remove(node *wNode) {
	nodePrev := node.prev
	nodeNext := node.next

	if node.prev != nil {
		(node.prev).next = nodeNext
	}

	if node.next != nil {
		(node.next).prev = nodePrev
	}
}

func iterateList(message string, list *wNode) {
	fmt.Println("---" + message + "---")
	tempNext := list
	for tempNext != nil {
		fmt.Println(tempNext.data)
		tempNext = tempNext.next
	}
	fmt.Println("---" + message + "---")
}

func main() {
	wList := &wNode{data: 2}

	wList = wList.PushHead(5)
	iterateList("PushHead test: push 5", wList)
	wList.PushTail(12)
	iterateList("PushTail test: push 12", wList)
	node, _ := wList.Find(12)
	fmt.Println("---Find test: find 12---")
	fmt.Println(node.data) // should be 12
	fmt.Println("---Find test: find 12---")
	wList.PushTail(11)
	wList.PushTail(12)
	wList.PushTail(13)
	wList.PushTail(14)
	findTwo, _ := node.Find(11)
	wList.Remove(findTwo)
	iterateList("Remove + Find test: start at 12, remove 11", wList)
	return
}
