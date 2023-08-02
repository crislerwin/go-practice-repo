# Stack

This folder contains implementations of a stack data structure using Go programming language. The stack follows the Last-In-First-Out (LIFO) principle, where the last element pushed onto the stack will be the first one to be popped off.

### Stack Array

In this implementation, the stack is represented using a fixed-size array. The top index variable keeps track of the topmost element in the stack. The push operation increments the top index and adds the element at that index. The pop operation retrieves the element at the top index and decrements the top index. The stack is considered empty when the top index is -1, and considered full when the top index is equal to the size of the array minus 1.

### Methods

- [x] push()
- [x] pop()
- [x] length()
- [x] peak()
- [x] empty()

## Stack Linked List

In this implementation, the stack is represented using a linked list.
Each node in the linked list contains the value of the element and a reference to the next node.
The top pointer points to the topmost node in the stack.
The push operation creates a new node with the given element and makes it the new top, updating the next reference to the previous top node.
The pop operation removes the topmost node and updates the top pointer to the next node. The stack is considered empty when the top pointer is null

- [x] push()
- [x] pop()
- [x] show()
- [x] len()
- [x] isEmpty()
- [x] peak()

## Stack Singly Linked List

In this implementation, the stack is represented using a singly linked list.
Singly linked list to implement a stack offers several advantages.
The primary benefit is efficient memory usage, as each node in the linked list only requires storage for the data value and a single pointer to the next node.
This makes it suitable for scenarios where dynamic memory allocation is preferred or when the size of the stack can vary during program execution.

- [x] Push(): Add an element to the top of the stack.
- [x] Pop(): Remove and return the element from the top of the stack.
- [x] Peek(): View the element at the top of the stack without removing it.
- [x] IsEmpty(): Check if the stack is empty.
- [x] Size(): Get the current number of elements in the stack.
