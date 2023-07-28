package stack

var stackArray []any

func stackPush(n any) {
	stackArray = append([]any{n}, stackArray...)
}

func stackLenght() int {
	return len(stackArray)
}

func stackPop() any {
	pop := stackArray[0]
	stackArray = stackArray[1:]
	return pop
}
