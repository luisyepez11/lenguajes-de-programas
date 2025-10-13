package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func (i *Interfaz) eliminar(id string) {
	var nuevaLista []Tarea
	bandera := false
	for _, t := range i.lista.Lista {
		if t.Id != id {
			nuevaLista = append(nuevaLista, t)

		} else {
			bandera = true
		}
	}
	i.lista.Lista = nuevaLista
	jsonData, err := json.MarshalIndent(i.lista, "", "")
	if err != nil {
		fmt.Println("Error al eliminar tarea")
		return
	}
	if bandera {
		err = os.WriteFile(i.archivo, jsonData, 0644)
		fmt.Println("Tarea eliminada con exito")
	} else {
		fmt.Println("Tarea no encontrada")
	}

}
