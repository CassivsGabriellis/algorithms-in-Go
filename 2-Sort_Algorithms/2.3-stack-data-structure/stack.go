package main

import "fmt"

// The node type
type Node struct {
    value interface{}
    previous *Node
}

// The stack data structure
type Stack struct {
    length  int
    head    *Node
}

// Function to create a new stack instance
func NewStack() *Stack {
    return &Stack{}
}

// The Push() method adds a element in the top of the stack
func (s *Stack) Push(item interface{}) {
    node := &Node{value: item} //create the new node
    s.length++ //increase the stack length
    if s.head == nil {//if empty, set the head to the new node
        s.head = node
        return
    }

    // If it's not empty, add the new node to the head of the stack
    // and point the head to the new node
    node.previous = s.head //poiting to the previous node
    s.head = node //setting the head for the new node
}

// the max() function returns the maximum value of two numbers
func max(a, b int) int {
    if a > b {
        return a
    }

    return b
}

// The Pop() method deletes and returns the new element
// on the top of the stack
func (s *Stack) Pop() interface{} {
    if s.head == nil { // If the stack is empty, return 'nil'
        return nil
    }
    
    // If not empty, the node on the top is removed,
    // updating the head field with the previous node
    s.length = max(0, s.length - 1)
    head := s.head
    s.head = head.previous

    return head.value
}

// The Peek() method returns the value of the node on the top
// of the stack, without deleting it
func (s *Stack) Peek() interface{} {
    if s.head == nil {
        return nil
    }

    return s.head.value
}

func main() {
    stack := NewStack()
    stack.Push(1)
    stack.Push(2)
    stack.Push(3)

    fmt.Println(stack.Peek())

    fmt.Println(stack.Pop())
    fmt.Println(stack.Pop())

    fmt.Println(stack.Peek())
}
