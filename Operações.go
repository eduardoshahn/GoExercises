package main

import (
	"fmt"
)

func main() {
	receiveValues()
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
