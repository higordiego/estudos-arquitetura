# Abstract Factory Method

# Introdução
Abstract Factory tem o intuito de criar famílias de objetos relacionados por meio de interfaces sem que a classe concreta seja especificada.

## Objetivo
 - Permite que possa se utilizar varios objetos relacionados por meio de factory method.
 - Permite relações de objetos com base em uma interface.


## Implementação
O abstract factory trabalha muito similar ao [Factory Method](../factory/factory.md), tem na composição dele varios Factory Method.

No exemplo abaixo vamos criar interface chamada AbstractFactory do tipo interface.

```go
type AbstractFactory interface {
    mulherPerfeita()
}
```

Criamos uma interface e um metodo chamado ```mulherPerfeita()```, uma forma básica para entender como funciona.

O metodo ```mulherPerfeita()``` é um metodo fábrica.

Vamos criar a mulher perfeita..

```go
package main

import "fmt"

// Mulher -
type Mulher struct {
	nacionalidade string
	corDosOlhos   string
}

// AbstractFactory -
type AbstractFactory interface {
	MulherPerfeita() Mulher
}

// BrasileiraFactory -
type BrasileiraFactory struct {
}

// AmericanaFactory -
type AmericanaFactory struct {
}

// MulherPerfeita - brasileira
func (b BrasileiraFactory) MulherPerfeita() Mulher {
	return Mulher{nacionalidade: "brasileira", corDosOlhos: "castanhos"}
}

// MulherPerfeita - americana
func (a AmericanaFactory) MulherPerfeita() Mulher {
	return Mulher{nacionalidade: "americana", corDosOlhos: "azul"}
}

// factory method
func pegarTipoDeMulher(typeMulher string) Mulher {

	var fabrica AbstractFactory
	switch typeMulher {
	case "americana":
		fabrica = AmericanaFactory{}
		return fabrica.MulherPerfeita()
	case "brasileira":
		fabrica = BrasileiraFactory{}
		return fabrica.MulherPerfeita()
	}
	return Mulher{}
}
func main() {

	tipoMulherAmericana := "americana"
	a := pegarTipoDeMulher(tipoMulherAmericana)
	fmt.Println("nacionalidade", a.nacionalidade)
	fmt.Println("olhos", a.corDosOlhos)

	tipoMulherBrasileira := "brasileira"
	b := pegarTipoDeMulher(tipoMulherBrasileira)
	fmt.Println("nacionalidade", b.nacionalidade)
	fmt.Println("olhos", b.corDosOlhos)

}

```





