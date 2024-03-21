package main

import (
	"fmt"
	"time"
)

func main() {
	receiveValues()
	fmt.Println("\n")
	showOneToTen()
	fmt.Println("\n")
	showSum()
	fmt.Println("\n")
	showPairs()
	fmt.Println("\n")
	fibonacciNumbers()
}

// Recebe o input do usuário e chama a função math() para trabalhar as variaveis a e b
func receiveValues() {
	var a, b int

	fmt.Printf("Digite um número inteiro: ")
	fmt.Scanln(&a)
	fmt.Printf("Digite outro número inteiro: ")
	fmt.Scanln(&b)

	math(a, b)
}

// Faz todos os calculos com as váriaveis a e b, e depois mostra para o usuário
func math(a, b int) {
	soma := a + b
	subtracao := a - b
	multiplicacao := a * b
	divisao := a / b
	restoDivisao := a % b
	fmt.Printf("A soma dos dois números é: %d\nA subtração dos dois numeros é: %d\nA multiplicação dos dois numeros é: %d\nA divisão dos dois numeros é: %d\nO resto da divisão dos dois numeros é: %d\n", soma, subtracao, multiplicacao, divisao, restoDivisao)

}

func showOneToTen() {
	for i := 1; i < 11; i++ {
		fmt.Println(i)
		time.Sleep(time.Second / 5)
	}
}

func showSum() {
	soma := 0
	for i := 1; i < 101; i++ {
		soma += i
	}
	fmt.Println(soma)
}

func showPairs() {
	for i := 1; i < 51; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}

func fibonacciNumbers() {
	result := 0
	var fibonacciList [19]int
	fibonacciList[0] = 1
	fibonacciList[1] = 1
	for i := 0; i < 17; i++ {
		result = fibonacciList[i] + fibonacciList[i + 1]
		fibonacciList[i+2] = result
	}
	fmt.Println(fibonacciList)
}
