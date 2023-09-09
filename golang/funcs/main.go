package main

import "fmt"

// Aula de Funcs

// Func normal
func soma(x int, y int) int {
	return x + y
}

// Func Variadica - Recebe um slice de inteiros

func somaTudo(x ...int) int {
	resultado := 0
	for _, v := range x {
		resultado += v
	}
	return resultado
}

type Pessoa struct {
	name string
	age  int
}

func (p Pessoa) andar() {
	println(p.name, "andou")
}

func (p Pessoa) falar() {
	println(p.name, "falou")
}

func (p Pessoa) apresentar() {
	println(p.name, "e", p.age, "anos")
}

func main() {

	result := func(x ...int) func() int {
		res := 0
		for _, v := range x {
			res += v
		}
		return func() int {
			return res
		}
	}

	person := Pessoa{
		name: "John",
		age:  30,
	}

	person.andar()
	person.falar()
	person.apresentar()
	println(result(1, 2, 3, 4, 5)())

	resultado := soma(3, 4)
	println(resultado)
	println(somaTudo(1, 2, 3, 4, 5))

	complexNum := 3 + 2i
	fmt.Println(complexNum)

}
