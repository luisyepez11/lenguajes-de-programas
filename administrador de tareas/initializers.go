package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

func (i *Interfaz) iniciarInterfaz() {
	data, err := os.ReadFile(i.archivo)
	if err != nil {
		fmt.Println("Archivo no encontrado, creando 'lista.json'")
		contenidoInicial := []byte(`{"lista":[]}`)
		os.WriteFile(i.archivo, contenidoInicial, 0644)
		return
	}

	err = json.Unmarshal(data, &i.lista)
	if err != nil {
		fmt.Println("Error al decodificar el JSON:", err)
		return
	}

}

func (i *Interfaz) listaComandos() {
	fmt.Println()
	expectedCommands := []string{"add: añade una tarea", "see: Muestra las tarea", "seeP: Muestra las tareaen proceso", "seeC: Muestra las tarea completadas", "seeI: Muestra las tarea que no se han realizado", "pop: Elimina una tarea", "p: Modificar porcentaje", "c: Completar Tarea", "n: Modificar nombre de tarea"}

	for _, t := range expectedCommands {
		fmt.Println(t)
	}
	println()
}

func (c *Cmd) inicia(interfaz Interfaz) {
	c.rootCmd = cobra.Command{Use: "app"}
	//comandos
	var cmdCom = &cobra.Command{
		Use:   "help ",
		Short: "Comando para ver todos los comandos",
		Long:  "Este comando te permite ver todos los comandos",
		Args:  cobra.MinimumNArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			interfaz.listaComandos()
		},
	}
	//añadir tarea
	var cmdAdd = &cobra.Command{
		Use:   "add [nombre]",
		Short: "Comando para añadit tareas",
		Long:  "Este comando agrega una tarea y la inicia en 0% y en estado incompleta",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			interfaz.guardar_tarea(args[0])
		},
	}
	//ver tareas
	var cmdSee = &cobra.Command{
		Use:   "see",
		Short: "Comando para ver tareas",
		Long:  "Este comando que te permite ver las tareas",
		Args:  cobra.MinimumNArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			interfaz.lista_tareas()
		},
	}
	//vre tareas completadas
	var cmdSeecompleted = &cobra.Command{
		Use:   "seeC",
		Short: "Comando para ver tareas Completadas",
		Long:  "Este comando que te permite ver las tareas Completadas",
		Args:  cobra.MinimumNArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			interfaz.listaTareasCompletadas()
		},
	}
	//ver tareas en proceso
	var cmdSeeProcess = &cobra.Command{
		Use:   "seeP",
		Short: "Comando para ver tareas en proceso",
		Long:  "Este comando que te permite ver las tareas en proceso",
		Args:  cobra.MinimumNArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			interfaz.listaTareasProcesos()
		},
	}
	//ver tareas incompletas
	var cmdInitiated = &cobra.Command{
		Use:   "seeI",
		Short: "Comando para ver tareas que no se han realizado",
		Long:  "Este comando que te permite ver las tareas que no se han realizado",
		Args:  cobra.MinimumNArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			interfaz.listaTareasIniciadas()
		},
	}
	//eliminar una tarea
	var cmdDelete = &cobra.Command{
		Use:   "pop [id]",
		Short: "Comando para borrar tareas",
		Long:  "Comando para eliminar permanentemente un registro de una tarea",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			interfaz.eliminar(args[0])
		},
	}
	//actualizar nombre
	var cmdNombre = &cobra.Command{
		Use:   "n [id] [nombre]",
		Short: "Comando para borrar tareas",
		Long:  "Comando para eliminar permanentemente un registro de una tarea",
		Args:  cobra.MinimumNArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			interfaz.actualizarNombre(args[0], args[1])
		},
	}
	//actualizar estado
	var cmdEstado = &cobra.Command{
		Use:   "c",
		Short: "Comando para borrar tareas",
		Long:  "Comando para eliminar permanentemente un registro de una tarea",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			interfaz.completar(args[0])
		},
	}
	//actualizar porcentaje
	var cmdPorcentaje = &cobra.Command{
		Use:   "p [id] [porcentaje]",
		Short: "Comando para borrar tareas",
		Long:  "Comando para eliminar permanentemente un registro de una tarea",
		Args:  cobra.MinimumNArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			porsetaje, err := strconv.ParseFloat(args[1], 32)
			if err != nil {
				fmt.Println(err)
			} else {
				interfaz.ModificarPorcentaje(args[0], porsetaje)
			}

		},
	}
	c.rootCmd.AddCommand(cmdAdd)
	c.rootCmd.AddCommand(cmdSee)
	c.rootCmd.AddCommand(cmdDelete)
	c.rootCmd.AddCommand(cmdPorcentaje)
	c.rootCmd.AddCommand(cmdEstado)
	c.rootCmd.AddCommand(cmdNombre)
	c.rootCmd.AddCommand(cmdCom)
	c.rootCmd.AddCommand(cmdSeeProcess)
	c.rootCmd.AddCommand(cmdSeecompleted)
	c.rootCmd.AddCommand(cmdInitiated)
	c.rootCmd.Execute()
}
