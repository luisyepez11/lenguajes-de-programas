package main

func main() {
	var tarea []Tarea
	lista := Lista{tarea}
	interfaz := Interfaz{"lista.json", lista}
	interfaz.iniciarInterfaz()
	sistema := Cmd{}
	sistema.inicia(interfaz)

}
