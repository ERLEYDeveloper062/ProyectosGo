package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"Proyecto1/automata"

	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

func pertenece(cadena string, automata *automata.Automata) bool {
	estadoActual := automata.EstadoInicial

	for _, entrada := range cadena {
		if _, existeTransicion := automata.Transiciones[estadoActual][string(entrada)]; existeTransicion {
			estadoActual = automata.Transiciones[estadoActual][string(entrada)]
			fmt.Printf("El estado actual es %s \r\n", estadoActual)
		} else {
			return false
		}
	}

	for _, estadoFinal := range automata.EstadosFinales {
		if estadoActual == estadoFinal {
			return true
		}
	}
	return false
}

func completarAutomata(automata *automata.Automata) {
	sumidero := fmt.Sprintf("S%d", len(automata.Estados))
	automata.Estados = append(automata.Estados, sumidero)

	//Completamos las transiciones faltantes con el sumidero
	for _, estado := range automata.Estados {
		if _, existeTransiscion := automata.Transiciones[estado]; !existeTransiscion {
			automata.Transiciones[estado] = make(map[string]string)
		}
		for _, entrada := range automata.Entradas {
			if _, existeTransiciones := automata.Transiciones[estado][entrada]; !existeTransiciones {
				automata.Transiciones[estado][entrada] = sumidero
			}
		}
	}

	//Agregamos transiciones desde el estado sumidero, asi como para las otras entradas
	automata.Transiciones[sumidero] = map[string]string{}
	for _, entrada := range automata.Entradas {
		automata.Transiciones[sumidero][entrada] = sumidero
	}
}

func automataCompleto(automata *automata.Automata) bool {
	for _, estado := range automata.Estados {
		for _, entrada := range automata.Entradas {
			if _, existeTransicion := automata.Transiciones[estado][entrada]; !existeTransicion {
				return false
			}
		}
	}

	fmt.Println("El automata está completo")
	return true
}

func cargarJSON(rutaArchivo string, datos interface{}) error {
	file, err := os.Open(rutaArchivo)
	fmt.Printf("entra a file %v\n\r", file)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}

	//automata := Automata{}
	err = json.Unmarshal(bytes, &datos)
	if err != nil {
		fmt.Println(err)
	}

	return nil
}

func main() {
	fmt.Println("Inicio del proyecto")

	myApp := app.New()
	myWindow := myApp.NewWindow("Hola")

	myWindow.SetContent(widget.NewLabel("Hola mundo"))
	myWindow.ShowAndRun()

	automata1 := automata.NewAutomata(
		[]string{"q1", "q2"},
		[]string{"0", "1"},
		map[string]map[string]string{
			"q1": {"0": "q1", "1": "q2"},
			"q2": {"0": "q2", "1": "q1"}},
		"q1",
		[]string{"q2"},
	)

	pertenece("0011", automata1)

	/*for llaveExterna, mapaInterno := range automata1.GetTransicion() {
		fmt.Println("Llave externa", llaveExterna)

		for llaveInterna, valor := range mapaInterno {
			fmt.Println("Llave interna", llaveInterna, "Valor: ", valor)
		}
	}*/

	/*print(automata1.GetEstadoInicial())
	print(automata1.GetTransicion())*/
}
