package main

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestIniciarInterfaz(t *testing.T) {

	testFile := "test_nonexistent.json"
	interfaz := Interfaz{archivo: testFile}

	defer os.Remove(testFile)

	interfaz.iniciarInterfaz()

	_, err := os.Stat(testFile)
	assert.Nil(t, err)

	data, err := os.ReadFile(testFile)
	assert.Nil(t, err)

	var lista Lista
	err = json.Unmarshal(data, &lista)
	assert.Nil(t, err)
	assert.Empty(t, lista.Lista)
}

func TestIniciarInterfaz_ArchivoExistente(t *testing.T) {
	testFile := "test_existing.json"
	initialData := Lista{
		Lista: []Tarea{
			{Id: "1", Nombre: "Tarea existente", Porcentaje: 50, Estado: EstadoAvanzado},
		},
	}

	data, err := json.Marshal(initialData)
	require.Nil(t, err)
	err = os.WriteFile(testFile, data, 0644)
	require.Nil(t, err)

	defer os.Remove(testFile)

	interfaz := Interfaz{archivo: testFile}
	interfaz.iniciarInterfaz()

	assert.Len(t, interfaz.lista.Lista, 1)
	assert.Equal(t, "Tarea existente", interfaz.lista.Lista[0].Nombre)
	assert.Equal(t, 50.0, interfaz.lista.Lista[0].Porcentaje)
}

func TestGuardarTarea(t *testing.T) {
	testFile := "test_guardar.json"
	interfaz := Interfaz{archivo: testFile}
	interfaz.lista = Lista{Lista: []Tarea{}}

	defer os.Remove(testFile)

	interfaz.guardar_tarea("Nueva tarea de prueba")

	assert.Len(t, interfaz.lista.Lista, 1)
	assert.Equal(t, "Nueva tarea de prueba", interfaz.lista.Lista[0].Nombre)
	assert.Equal(t, 0.0, interfaz.lista.Lista[0].Porcentaje)
	assert.Equal(t, EstadoIncompleto, interfaz.lista.Lista[0].Estado)

	data, err := os.ReadFile(testFile)
	assert.Nil(t, err)

	var listaGuardada Lista
	err = json.Unmarshal(data, &listaGuardada)
	assert.Nil(t, err)
	assert.Len(t, listaGuardada.Lista, 1)
	assert.Equal(t, "Nueva tarea de prueba", listaGuardada.Lista[0].Nombre)
}

func TestEliminarTarea(t *testing.T) {
	testFile := "test_eliminar.json"
	tareaID := uuid.New().String()

	interfaz := Interfaz{archivo: testFile}
	interfaz.lista = Lista{
		Lista: []Tarea{
			{Id: tareaID, Nombre: "Tarea a eliminar", Porcentaje: 0, Estado: EstadoIncompleto},
			{Id: "otra-tarea", Nombre: "Otra tarea", Porcentaje: 50, Estado: EstadoAvanzado},
		},
	}

	initialData, err := json.Marshal(interfaz.lista)
	require.Nil(t, err)
	err = os.WriteFile(testFile, initialData, 0644)
	require.Nil(t, err)

	defer os.Remove(testFile)

	interfaz.eliminar(tareaID)

	assert.Len(t, interfaz.lista.Lista, 1)
	assert.Equal(t, "Otra tarea", interfaz.lista.Lista[0].Nombre)

	data, err := os.ReadFile(testFile)
	assert.Nil(t, err)

	var listaActualizada Lista
	err = json.Unmarshal(data, &listaActualizada)
	assert.Nil(t, err)
	assert.Len(t, listaActualizada.Lista, 1)
	assert.Equal(t, "Otra tarea", listaActualizada.Lista[0].Nombre)
}

func TestActualizarNombre(t *testing.T) {
	testFile := "test_actualizar_nombre.json"
	tareaID := uuid.New().String()

	interfaz := Interfaz{archivo: testFile}
	interfaz.lista = Lista{
		Lista: []Tarea{
			{Id: tareaID, Nombre: "Nombre antiguo", Porcentaje: 0, Estado: EstadoIncompleto},
		},
	}

	initialData, err := json.Marshal(interfaz.lista)
	require.Nil(t, err)
	err = os.WriteFile(testFile, initialData, 0644)
	require.Nil(t, err)

	defer os.Remove(testFile)

	interfaz.actualizarNombre(tareaID, "Nuevo nombre")

	assert.Len(t, interfaz.lista.Lista, 1)
	assert.Equal(t, "Nuevo nombre", interfaz.lista.Lista[0].Nombre)

	data, err := os.ReadFile(testFile)
	assert.Nil(t, err)

	var listaActualizada Lista
	err = json.Unmarshal(data, &listaActualizada)
	assert.Nil(t, err)
	assert.Equal(t, "Nuevo nombre", listaActualizada.Lista[0].Nombre)
}

func TestCompletarTareaInterfaz(t *testing.T) {
	testFile := "test_completar.json"
	tareaID := uuid.New().String()

	interfaz := Interfaz{archivo: testFile}
	interfaz.lista = Lista{
		Lista: []Tarea{
			{Id: tareaID, Nombre: "Tarea por completar", Porcentaje: 50, Estado: EstadoAvanzado},
		},
	}

	initialData, err := json.Marshal(interfaz.lista)
	require.Nil(t, err)
	err = os.WriteFile(testFile, initialData, 0644)
	require.Nil(t, err)

	defer os.Remove(testFile)

	interfaz.completar(tareaID)

	assert.Len(t, interfaz.lista.Lista, 1)
	assert.Equal(t, 100.0, interfaz.lista.Lista[0].Porcentaje)
	assert.Equal(t, EstadoCompletado, interfaz.lista.Lista[0].Estado)

	data, err := os.ReadFile(testFile)
	assert.Nil(t, err)

	var listaActualizada Lista
	err = json.Unmarshal(data, &listaActualizada)
	assert.Nil(t, err)
	assert.Equal(t, 100.0, listaActualizada.Lista[0].Porcentaje)
	assert.Equal(t, EstadoCompletado, listaActualizada.Lista[0].Estado)
}

func TestModificarPorcentaje(t *testing.T) {
	testFile := "test_modificar_porcentaje.json"
	tareaID := uuid.New().String()

	interfaz := Interfaz{archivo: testFile}
	interfaz.lista = Lista{
		Lista: []Tarea{
			{Id: tareaID, Nombre: "Tarea test", Porcentaje: 0, Estado: EstadoIncompleto},
		},
	}

	initialData, err := json.Marshal(interfaz.lista)
	require.Nil(t, err)
	err = os.WriteFile(testFile, initialData, 0644)
	require.Nil(t, err)

	defer os.Remove(testFile)

	interfaz.ModificarPorcentaje(tareaID, 75.0)

	assert.Len(t, interfaz.lista.Lista, 1)
	assert.Equal(t, 75.0, interfaz.lista.Lista[0].Porcentaje)
	assert.Equal(t, EstadoAvanzado, interfaz.lista.Lista[0].Estado)

	data, err := os.ReadFile(testFile)
	assert.Nil(t, err)

	var listaActualizada Lista
	err = json.Unmarshal(data, &listaActualizada)
	assert.Nil(t, err)
	assert.Equal(t, 75.0, listaActualizada.Lista[0].Porcentaje)
	assert.Equal(t, EstadoAvanzado, listaActualizada.Lista[0].Estado)
}
