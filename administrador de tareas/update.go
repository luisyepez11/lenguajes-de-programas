package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

// actualizar nombre
func (t *Tarea) editarNombreTarea(nuevoNombre string) (string, error) {
	if nuevoNombre != "" {
		t.Nombre = nuevoNombre
		fmt.Println("\nNombre Modificado")
		t.mostrarTarea()
		fmt.Println()
		return nuevoNombre, nil
	} else if nuevoNombre == "" {
		return "", errors.New("No ingreso ningun nombre")
	} else {
		return "", errors.New("Error al actualizar el nombre")
	}
}

func (i *Interfaz) actualizarNombre(id string, nombre string) {
	var nuevaLista []Tarea
	bandera := false
	for _, t := range i.lista.Lista {
		if t.Id == id {
			t.editarNombreTarea(nombre)
			nuevaLista = append(nuevaLista, t)
			bandera = true
		} else {
			nuevaLista = append(nuevaLista, t)
		}
	}
	i.lista.Lista = nuevaLista
	jsonData, err := json.MarshalIndent(i.lista, "", "")
	if err != nil {
		return
	}
	if bandera {
		err = os.WriteFile(i.archivo, jsonData, 0644)
		fmt.Println("Nombre actualizado con exito")
	} else {
		fmt.Println("Tarea no encontrada")
	}

}
//actualizar porcentaje
func (t *Tarea) actualizarEstadoPorPorcentaje() {
	switch {
	case t.Porcentaje == 0:
		t.Estado = EstadoIncompleto
	case t.Porcentaje > 0 && t.Porcentaje < 50:
		t.Estado = EstadoProceso
	case t.Porcentaje >= 50 && t.Porcentaje < 100:
		t.Estado = EstadoAvanzado
	case t.Porcentaje == 100:
		t.Estado = EstadoCompletado
	default:
		t.Estado = EstadoIncompleto
	}
}
func (t *Tarea) editarPorcentajeTarea(porcentaje float64) (float64, error) {
	if porcentaje <= 100 {
		t.Porcentaje = porcentaje
		fmt.Println("\n Porcentaje Modificado ")
		t.actualizarEstadoPorPorcentaje()
		t.mostrarTarea()
		return porcentaje, nil
	} else {
		fmt.Println("Error el porcentaje debe ser menor o igual a 100")
		return -1, errors.New("Error el porcentaje debe ser menor o igual a 100")
	}

}

func (i *Interfaz) ModificarPorcentaje(id string, porcentaje float64) {
	var nuevaLista []Tarea
	bandera := false
	for _, t := range i.lista.Lista {
		if t.Id == id {
			t.editarPorcentajeTarea(porcentaje)
			nuevaLista = append(nuevaLista, t)
			bandera = true
		} else {
			nuevaLista = append(nuevaLista, t)
		}
	}
	i.lista.Lista = nuevaLista
	jsonData, err := json.MarshalIndent(i.lista, "", "")
	if err != nil {
		return
	}
	if bandera {
		err = os.WriteFile(i.archivo, jsonData, 0644)

	} else {
		fmt.Println("Tarea no encontrada")
	}
}
//actualizar  a estado completado
func (t *Tarea) CompletarTarea() {
	t.Porcentaje = 100
	t.Estado = EstadoCompletado
	fmt.Println("\nTarea Completada")
	t.actualizarEstadoPorPorcentaje()
	t.mostrarTarea()
	fmt.Println()
}

func (i *Interfaz) completar(id string) {
	var nuevaLista []Tarea
	bandera := false
	for _, t := range i.lista.Lista {
		if t.Id == id {
			t.CompletarTarea()
			bandera = true
			nuevaLista = append(nuevaLista, t)
		} else {
			nuevaLista = append(nuevaLista, t)
		}
	}
	i.lista.Lista = nuevaLista
	jsonData, err := json.MarshalIndent(i.lista, "", "")
	if err != nil {
		return
	}
	if bandera {
		err = os.WriteFile(i.archivo, jsonData, 0644)
		fmt.Println("Tarea completada")
	} else {
		fmt.Println("Tarea no encontrada")
	}

}
