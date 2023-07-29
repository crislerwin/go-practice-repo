package stack

var stackArray []any

func stackPush(n any) {
	stackArray = append([]any{n}, stackArray...)
}

func stackLength() int {
	return len(stackArray)
}

func stackPop() any {
	pop := stackArray[0]
	stackArray = stackArray[1:]
	return pop
}

func stackPeak() any {
	return stackArray[0]
}
