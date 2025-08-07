package main

import "fmt"

type Node struct {
	Val  int
	Next *Node
}

var head *Node = nil
var tail *Node = nil

func createNode(val int) *Node {
	newNode := new(Node)
	newNode.Val = val
	return newNode
}
func insertAtHead(val int) {
	newNode := createNode(val)
	newNode.Val = val
	newNode.Next = head
	head = newNode
	if tail == nil {
		tail = newNode
	}
}

func print_ll() {
	newNode := head
	for newNode != nil {
		if newNode.Next != nil {
			fmt.Printf("%d -> ", newNode.Val)
			newNode = newNode.Next
		} else {
			fmt.Printf("%d", newNode.Val)
			newNode = newNode.Next
		}
	}
}

func main() {
	insertAtHead(1)
	insertAtHead(2)
	insertAtHead(3)
	insertAtHead(7)
	print_ll()

}
