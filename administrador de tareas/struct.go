package main

import (
	"time"

	"github.com/spf13/cobra"
)

type Estado string

const (
	EstadoIncompleto Estado = "Incompleto"
	EstadoProceso    Estado = "Proceso"
	EstadoAvanzado   Estado = "Avanzado"
	EstadoCompletado Estado = "Completado"
)

type Tarea struct {
	Id         string    `json:"id"`
	Nombre     string    `json:"nombre"`
	Porcentaje float64   `json:"porcentaje"`
	Estado     Estado    `json:"estado"`
	Creacion   time.Time `json:"fecha_creacion"`
}

type Interfaz struct {
	archivo string
	lista   Lista
}

type Lista struct {
	Lista []Tarea `json:"lista"`
}

type Cmd struct {
	rootCmd cobra.Command
}
