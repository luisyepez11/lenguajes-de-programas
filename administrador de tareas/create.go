package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"
)

func (i *Interfaz) guardar_tarea(nombre string) {
	id := uuid.New()
	fmt.Println(id.String())
	ahora := time.Now()
	año, mes, dia := ahora.Date()
	fechaCreacion := time.Date(año, mes, dia, 0, 0, 0, 0, ahora.Location())
	tarea := Tarea{id.String(), nombre, 0, EstadoIncompleto, fechaCreacion}
	i.lista.Lista = append(i.lista.Lista, tarea)
	fmt.Println("tarea agregada exitosamente ")
	jsonData, err := json.MarshalIndent(i.lista, "", "  ")
	if err != nil {
		return
	}
	err = os.WriteFile(i.archivo, jsonData, 0644)
}
