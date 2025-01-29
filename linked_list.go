package main

import "fmt"

type Node struct {
	Data interface{}
	Next *Node
}

type LinkedList struct {
	head *Node
}

//insert element into linked_list
/*
	check if head is not empty,
	create new node
	if head empty insert new node into head
	if head !empty
		find next empty node
		insert into non empty node
*/
func (ll *LinkedList) InsertNode(data interface{}) {
	new_node := &Node{Data: data}
	if ll.head == nil {
		ll.head = new_node
	} else {
		//set the current node as head
		current_node := ll.head
		//keep looping to get to end i.e an empty node or if you;ve gotten to the last node that the next is nil
		for current_node.Next != nil {
			current_node = current_node.Next
		}
		//set the node's next to the new node
		current_node.Next = new_node
	}
}

func (ll *LinkedList) Print() {
	//set the first node to head
	current_node := ll.head
	for current_node.Next != nil {
		fmt.Printf("%v ->", current_node.Data)
		current_node = current_node.Next
	}
	fmt.Print("end")
}
