package math

import "fmt"

func IMC(weight, height float32) string {
	var result = weight / (height * height)

	if result < 18.5 {
		return fmt.Sprint("Seu IMC é: ", result, " portando você está abaixo do peso ideal.")
	} else if result >= 18.5 && result < 25 {
		return fmt.Sprint("Seu IMC é: ", result, " portando você está no peso ideal.")
	} else if result >= 25 && result < 30 {
		return fmt.Sprint("Seu IMC é: ", result, " portando você está acima do peso ideal.")
	} else if result >= 30 && result < 35 {
		return fmt.Sprint("Seu IMC é: ", result, " portando você está com Obesidade de Grau I.")
	} else if result >= 35 && result < 40 {
		return fmt.Sprint("Seu IMC é: ", result, " portando você está com Obesidade de Grau II.")
	} else {
		return fmt.Sprint("Seu IMC é: ", result, " portando você está com Obesidade Mórbida.")
	}
}
