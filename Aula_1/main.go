/*
Na serie, "Praticando Golang", vamos criar pequenos processos e funções, e vamos aplicar
os testes em cada uma deles, para irmos testando e aprendendo uma das mais importantes
ação de um programador, que é "Fazer os testes individuais de cada função"
*/
package main

import "fmt"

//Vamos criar uma conversor de "Milhas para KM, e talbem o contrario KM para Milhas"
// Usaremos a nimenclatura 'INGLES' para o codigo.

const miles float32 = 1609.0 // 1 milhas = 1.609 metros.

func milesConvertToKM(m float32) float32 {
	result := m * miles
	fmt.Println(m, "MIlhas convertidas em KM é:", result, "metros")
	return result
}

func kmConvertToMiles(km float32) {
	result := (km / miles) * 1000
	fmt.Println(km, "Kilometros convertidos em milhas é:", result)
}

func main() {
	milesConvertToKM(3)
	kmConvertToMiles(3)

}
