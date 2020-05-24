package main

import "fmt"

type construtora interface {
	inserirJanelasTipo()
	inserirPortasTipo()
	pegarCasa() casa
}

func pegarConstrutora(casaTipo string) construtora {
	if casaTipo == "normal" {
		return &normalCasa{}
	}
	if casaTipo == "linda" {
		return &lindaCasa{}
	}
	return nil
}

type normalCasa struct {
	janelaTipo string
	portaTipo  string
}

func newNormalConstrutora() *normalCasa {
	return &normalCasa{}
}

func (b *normalCasa) inserirJanelasTipo() {
	b.janelaTipo = "Janela Normal"
}

func (b *normalCasa) inserirPortasTipo() {
	b.portaTipo = "Porta normal"
}

func (b *normalCasa) pegarCasa() casa {
	return casa{
		janelaTipo: b.janelaTipo,
		portaTipo:  b.portaTipo,
	}
}

type lindaCasa struct {
	janelaTipo string
	portaTipo  string
}

func newLindaConstrutora() *lindaCasa {
	return &lindaCasa{}
}

func (b *lindaCasa) inserirJanelasTipo() {
	b.janelaTipo = "Janela Normal"
}

func (b *lindaCasa) inserirPortasTipo() {
	b.portaTipo = "Porta normal"
}

func (b *lindaCasa) pegarCasa() casa {
	return casa{
		janelaTipo: b.janelaTipo,
		portaTipo:  b.portaTipo,
	}
}

type casa struct {
	janelaTipo string
	portaTipo  string
}

type diretor struct {
	construtora construtora
}

func novoDiretor(b construtora) *diretor {
	return &diretor{
		construtora: b,
	}
}

func (d *diretor) inserirConstrutora(c construtora) {
	d.construtora = c
}

func (d diretor) construirCasa() casa {
	d.construtora.inserirJanelasTipo()
	d.construtora.inserirPortasTipo()
	return d.construtora.pegarCasa()
}

func main() {
	normalBuild := pegarConstrutora("normal")

	diretor := novoDiretor(normalBuild)

	normalCasa := diretor.construirCasa()

	fmt.Println("Normal casa tipo de janela: ", normalCasa.janelaTipo)
	fmt.Println("Normal casa tipo de porta: ", normalCasa.portaTipo)

}
