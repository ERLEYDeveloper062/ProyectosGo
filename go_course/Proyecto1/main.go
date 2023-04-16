package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"

	"Proyecto1/automata"

	"fyne.io/fyne/app"

	//"fyne.io/fyne/container"
	"fyne.io/fyne/layout"
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
	fmt.Printf("El sumidero es: %s \r\n", sumidero)

	// Completamos las transiciones faltantes con el sumidero
	for _, estado := range automata.Estados {
		// Si el estado es una cadena vacía, lo reemplazamos por el sumidero
		if estado == "" {
			estado = sumidero
		}
		if _, existeTransicion := automata.Transiciones[estado]; !existeTransicion {
			automata.Transiciones[estado] = make(map[string]string)
		}
		for _, entrada := range automata.Entradas {
			if _, existeTransicion := automata.Transiciones[estado][entrada]; !existeTransicion {
				// Si la transición apunta a una cadena vacía, la reemplazamos por el sumidero
				if automata.Transiciones[estado][entrada] == "" {
					automata.Transiciones[estado][entrada] = sumidero
				} else {
					automata.Transiciones[estado][entrada] = automata.Transiciones[estado][entrada]
				}
			}
		}
	}

	// Reemplazamos cualquier estado final que sea una cadena vacía por el sumidero
	for i, estadoFinal := range automata.EstadosFinales {
		if estadoFinal == "" {
			automata.EstadosFinales[i] = sumidero
		}
	}

	// Agregamos transiciones desde el estado sumidero, así como para las otras entradas
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

func interfazGrafica(automata *automata.Automata) {

	a := app.New()
	w := a.NewWindow("ACME")

	estadosWidget := widget.NewLabel("Estados: " + strings.Join(automata.GetEstados(), ", "))
	entradasWidget := widget.NewLabel("Entradas: " + strings.Join(automata.GetEntradas(), ", "))
	/*transicionesWidget := widget.NewLabel("Transiciones: " + strings.Join(automata.Transiciones.estado.entrada, ", ") + "\n")

	for i := 1; i < len(transiciones); i++ {
		transicionesWidget.SetText(transicionesWidget.Text + strings.Join(transiciones[i], ", ") + "\n")
	}*/

	estadoInicialWidget := widget.NewLabel("Estado Inicial: " + automata.GetEstadoInicial())
	estadosFinalesWidget := widget.NewLabel("Estados Finales: " + strings.Join(automata.EstadosFinales, ", "))

	cadenaEntry := widget.NewEntry()
	evaluarButton := widget.NewButton("Evaluar", func() {
		cadena := cadenaEntry.Text
		if pertenece(cadena, automata) {
			w.SetContent(widget.NewVBox(
				widget.NewLabel(fmt.Sprintf("La cadena '%s' es aceptada por el automata", cadena)),
				layout.NewSpacer(),
				widget.NewButton("Cerrar", func() {
					a.Quit()
				}),
			))
		} else {
			w.SetContent(widget.NewVBox(
				widget.NewLabel(fmt.Sprintf("La cadena '%s' no es aceptada por el automata", cadena)),

				layout.NewSpacer(),
				widget.NewButton("Cerrar", func() {
					a.Quit()
				}),
			))
		}
	})

	content := widget.NewVBox(
		//container.NewVBox(Opcion, Opcion2, lector, respuesta1, respuesta2, respuesta3, respuesta4, respuesta5),
		//content.Resize(fyne.NewSize(280, 0)),
		//content.Move(fyne.NewPos(10, 100)),
		//content5 := container.NewWithoutLayout(banner, content)
		//content4 := container.NewWithoutLayout(content5, content2, content3, content6, content8, contenedor9)
		estadosWidget,
		entradasWidget,
		layout.NewSpacer(),
		estadoInicialWidget,
		estadosFinalesWidget,
		cadenaEntry,
		evaluarButton,
	)

	w.SetContent(content)
	w.ShowAndRun()
}

func main() {
	fmt.Println("Inicio del proyecto")

	var automata1 automata.Automata
	err := cargarJSON("automata1.json", &automata1)
	if err != nil {
		panic(err)
	}

	// Si el autómata no está completo, lo completamos
	if !automataCompleto(&automata1) {
		fmt.Println("El automata está incompeto lo completamos")
		completarAutomata(&automata1)
	}

	fmt.Printf("estado Inicial %v\n\r", automata1.GetEstadoInicial())
	fmt.Printf("estado finales %v\n\r", automata1.EstadosFinales)
	fmt.Printf("estados %v\n\r", automata1.Estados)
	fmt.Printf("entradas %v\n\r", automata1.Entradas)
	fmt.Printf("transiciones %v\n\r", automata1.Transiciones)

	interfazGrafica(&automata1)

}
