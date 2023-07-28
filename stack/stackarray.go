package stack

var stackArray []any

func stackPush(n any) {
	stackArray = append([]any{n}, stackArray...)
}
