package automata

import (
	"errors"
	"fmt"
)

type Automata struct {
	Estados        []string                     `json:"estados"`
	Entradas       []string                     `json:"entradas"`
	Transiciones   map[string]map[string]string `json:"transiciones"`
	EstadoInicial  string                       `json:"estadoInicial"`
	EstadosFinales []string                     `json:"estadosFinales"`
}

func NewAutomata(estados []string, entradas []string, transiciones map[string]map[string]string, estadoInicial string, estadosFinales []string) *Automata {
	return &Automata{
		Estados:        estados,
		Entradas:       entradas,
		Transiciones:   transiciones,
		EstadoInicial:  estadoInicial,
		EstadosFinales: estadosFinales,
	}
}

// Getters
func (automata *Automata) GetEstados() []string {
	return automata.Estados
}

func (automata *Automata) GetEntradas() []string {
	return automata.Entradas
}

func (automata *Automata) GetTransicion() map[string]map[string]string {
	return automata.Transiciones
}

func (automata *Automata) GetEstadoInicial() string {
	return automata.EstadoInicial
}

func (automata *Automata) GetEstadosFinales() []string {
	return automata.EstadosFinales
}

//Setters

func (automata *Automata) SetEstados(estados []string) {
	automata.Estados = estados
}

func (automata *Automata) SetEntradas(entradas []string) {
	automata.Entradas = entradas
}

func (automata *Automata) SetTransiciones(trancicion map[string]map[string]string) error {
	if trancicion == nil {
		return errors.New("El mapa de transicion no puede estar vacio")
	}

	for llaveExterna, mapaInterno := range trancicion {
		fmt.Println("LLave externa", llaveExterna)

		/*ok := automata.Estados[llaveExterna]
		if !ok{
			return fmt.Errorf("La clave de estado de origen %v no es valida", llaveExterna)
		}*/

		for llaveInterna, valor := range mapaInterno {
			fmt.Println("Llave interna: ", llaveInterna, "Valor: ", valor)
		}
	}
	automata.Transiciones = trancicion
	return nil
}

func (automata *Automata) SetEstadoInicial(estadoInicial string) {
	automata.EstadoInicial = estadoInicial
}

func (automata *Automata) SetEstadosFinales(estadoFinales []string) {
	automata.EstadosFinales = estadoFinales
}
