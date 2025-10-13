package main

import (
	"fmt"
	"strings"
)

// tareas
func (t *Tarea) mostrarTarea() {
	fmt.Println("Nombre de la tarea : " + t.Nombre + " Estado : " + string(t.Estado))
	fmt.Println("Su id es : " + t.Id)
	fmt.Println("Porcentaje completado:")
	num := t.Porcentaje * 2
	numeral := "#"
	sumador := "["
	for i := 0; i < int(num/10); i++ {
		sumador += numeral
	}
	numeral = "_"
	if num < 200 {
		for i := 0; i < (20 - int((num)/10)); i++ {
			sumador += numeral
		}
	}
	sumador += "]"
	fmt.Printf("%s %.2f%%", sumador, t.Porcentaje)
	fmt.Print("\nfecha de creacion : " + strings.Split(t.Creacion.String(), " ")[0])
	fmt.Println("\n------------------------------------------------")

}

func (i *Interfaz) lista_tareas() {
	fmt.Println()
	if len(i.lista.Lista) > 0 {
		for _, t := range i.lista.Lista {
			t.mostrarTarea()
		}
	} else {
		println("No hay tareas guardadas")
	}
	println()
}

// tareas completadas
func (i *Interfaz) listaTareasCompletadas() {
	fmt.Println()
	bandera := false
	if len(i.lista.Lista) > 0 {
		for _, t := range i.lista.Lista {
			if t.Estado == EstadoCompletado {
				t.mostrarTarea()
				bandera = true
			}

		}
		if !bandera {
			println("No hay tareas Completadas")
		}
	} else {
		println("No hay tareas guardadas")
	}
	println()
}

// tareas en proceso
func (i *Interfaz) listaTareasProcesos() {
	fmt.Println()
	bandera := false
	if len(i.lista.Lista) > 0 {
		for _, t := range i.lista.Lista {
			if t.Estado != EstadoCompletado && t.Estado != EstadoIncompleto {
				t.mostrarTarea()
				bandera = true
			}

		}
		if !bandera {
			println("No hay tareas guardadas")
		}
	} else {
		println("No hay tareas guardadas")
	}
	println()
}

// tareas sin progreso
func (i *Interfaz) listaTareasIniciadas() {
	bandera := false
	fmt.Println()
	if len(i.lista.Lista) > 0 {
		for _, t := range i.lista.Lista {
			if t.Estado == EstadoIncompleto {
				t.mostrarTarea()
				bandera = true
			}

		}
		if !bandera {
			println("No hay tareas guardadas")
		}
	} else {
		println("No hay tareas guardadas")
	}
	println()
}
