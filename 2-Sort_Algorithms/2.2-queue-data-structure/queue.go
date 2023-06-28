package main

import "fmt"

//The Node<T> type
type Node struct {
    value interface{}
    next *Node
}

//The Queue structure that points to the Node<T>
type Queue struct {
    length int
    head *Node
    tail *Node
}

//The "constructor" of the Queue
func NewQueue() *Queue {
    return &Queue{}
}

//Adding the new element in the queue
func (q *Queue) Enqueue(item interface{}) {
    node := &Node{value: item}
    q.length++;
    if q.tail == nil {
        q.tail = node
        q.head = node
        return
    }
    
    q.tail.next = node
    q.tail = node
}

//Popping the first element of the queue
func (q *Queue) Dequeue() interface{} {
    if q.head == nil {
        return nil
    }

    q.length--
    head := q.head
    q.head = q.head.next
    
    if q.head == nil {
        q.tail = nil
    }

    return head.value
}

//Return the current value of the head of the queue
func (q *Queue) Peek() interface{} {
    if q.head == nil {
        return nil
    }

    return q.head.value
}

func main() {
    queue := NewQueue()
    queue.Enqueue(1)
    queue.Enqueue(2)
    queue.Enqueue(3)
    
    fmt.Println(queue.Peek())

    fmt.Println(queue.Dequeue())
    fmt.Println(queue.Dequeue())

    fmt.Println(queue.Peek())
}
