package main

import (
	"fmt"
	"log"
)

type Stack struct {
	Items []interface{}
}

func (s *Stack) Push(item interface{}) {
	s.Items = append(s.Items, item)
}

func (s *Stack) Pop() {
	if len(s.Items) == 0 {
		log.Println("stack is empty")
	}
	last_insert_id := len(s.Items) - 1
	s.Items = s.Items[:last_insert_id]
}

func (s *Stack) Print() {
	fmt.Println(s.Items...)
}
