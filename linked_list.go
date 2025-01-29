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
func (ll *LinkedList) Insert(data interface{}) {
	//create a new node
	new_node := &Node{Data: data}
	if ll.head == nil {
		ll.head = new_node
	} else {
		//start from  the head
		current_node := ll.head
		for current_node.Next != nil {
			//set current node to the next one until you get to the end(empty node)
			current_node = current_node.Next
		}
		current_node.Next = new_node
	}
}

func (ll *LinkedList) PrintList() {
	current_node := ll.head
	for current_node != nil {
		fmt.Printf("%v ->", current_node.Data)
		current_node = current_node.Next
	}
	fmt.Println("end")
}
