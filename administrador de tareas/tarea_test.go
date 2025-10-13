package main

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestEditarNombreTarea(t *testing.T) {
	id := uuid.New()
	tarea_testing := Tarea{Id: id.String(), Nombre: "tarea1", Porcentaje: 0, Estado: EstadoIncompleto}
	result, err := tarea_testing.editarNombreTarea("nueva tarea")
	assert.Nil(t, err)
	assert.Equal(t, "nueva tarea", result)
}

func TestEditarPorcentajeTarea(t *testing.T) {
	id := uuid.New()
	tarea_testing := Tarea{Id: id.String(), Nombre: "tarea1", Porcentaje: 0, Estado: EstadoIncompleto}
	result, err := tarea_testing.editarPorcentajeTarea(50.00)
	assert.Nil(t, err)
	assert.Equal(t, 50.00, result)
}

func TestCompletarTarea(t *testing.T) {
	id := uuid.New()
	tarea_testing := Tarea{Id: id.String(), Nombre: "tarea1", Porcentaje: 0, Estado: EstadoIncompleto}
	tarea_testing.CompletarTarea()
	assert.Equal(t, 100.00, tarea_testing.Porcentaje)
	assert.Equal(t, EstadoCompletado, tarea_testing.Estado)
}
