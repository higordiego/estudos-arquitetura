# Factory Method

## Introdução
```O Factory Method é um padrão de design usado para definir uma interface de tempo de execução para criar um objeto. É chamado de fábrica porque cria vários tipos de objetos sem necessariamente saber que tipo de objeto ele cria ou como criá-lo.```

## Objetivo
 - Permite que as subclasses escolham o tipo de objeto a serem criados em tempo de execurção.
 - Ele fornece uma maneira simples de estender a familia de objetos com pequenas alterações no código do aplicativo.
 - Acomplamento flexível, eliminando a necessidade de vincular estruturas especificas do código.

 # Implementação

 O factory method define uma interface para criar objetos, mas permitem que os outros decidam quais objetos vai instanciar.

 No exemplo abaixo iremos utilizar uma fabrica de armas cuja a fabricação será de uma ak47 e uma maverick.
 

Criando o produtos da fabrica de armas

 ```go
type produto interface {
    setName(name string)
    setPower(name int)
    getName() string
    getPower() int
}
 ```

Criando uma estrutura e seus métodos de retorno da arma.

```go
type arma struct {
    name string
    power int
}

func (g *gun) setName(name string) {
    g.name = name
}

func (g *gun) getName() string {
    return g.name
}

func (g *gun) setPower(power int) {
    g.power = power
}

func (g *gun) getPower() int {
    return g.power
}
 ```

Criando os tipos ak47 e maverick
```go
type ak47 struct {
    arma
}

func newAk47() arma {
    return &ak47 {
        gun: gun{
            name: "ak47 arma",
            power: 4
        },
    }
}

type maverick struct {
    gun
}

func newMaveric()arma {
    return &maverick {
        gun: gun{
            name: "Maverick gun",
            power: 5
        },
    }
}
```

Criaremos a fabrica.

```go
import "fmt"
func getGun(tipoArma string) (arma, error) {
    if tipoArma == "ak47" {
        return newAk47(), nil
    }
    if tipoArma == "maverick" {
        return newMaverick(), nil
    }
     return nil, fmt.Errorf("Não fabricamos esse tipo de arma!")
}
```

Executando na função main()
```go
func main() {   
    ak47, _ : getGun("ak47")
    maverick, _ : getGun("maverick")

    imprimirDetalhes(ak47)
    imprimirDetalhes(maverick)

    func imprimirDetalhes(g arma) {
        fmt.Println("Gun: %v", g.getName())
        fmt.Println("Power %v", g.getName())
    }
}
```




