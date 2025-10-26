
use crate::lista::Lista;
use std::fs;
use std::io::{Read, Write};

use crate::crud::{create, read, update, delete};

pub struct Interfaz {
    pub lista: Lista,
    pub archivo: String,
}

impl Interfaz {
    pub fn new(archivo: String, lista: Lista) -> Self {
        Interfaz { lista, archivo }
    }

    pub fn crear_tarea(&mut self, nombre: String, porcentaje: f32, creacion: String) {

        create::crear_tarea(&mut self.lista, nombre, porcentaje,creacion);
    }

    pub fn ver_tareas(&self) {
  
        read::ver_tareas(&self.lista);
    }
    pub fn ver_tareas_completadas(&self) {
  
        read::ver_tareas_completadas(&self.lista);
    }
    pub fn ver_tareas_no_iniciadas(&self) {
  
        read::ver_tareas_no_iniciadas(&self.lista);
    }
    
    pub fn actualizar_tarea_nombre(&mut self, id: String, nombre: String) -> bool {
   
        update::actualizar_tarea_nombre(&mut self.lista, id, nombre)
    }

    pub fn actualizar_tarea_porcentaje(&mut self, id: String, porcentaje: f32) -> bool {

        update::actualizar_tarea_porcentaje(&mut self.lista, id, porcentaje)
    }

    pub fn actualizar_tarea_porcentaje_completa(&mut self, id: String) -> bool {

        update::actualizar_tarea_porcentaje_completa(&mut self.lista, id)
    }

    pub fn eliminar_tarea(&mut self, id: String) -> bool {

        delete::elimminar_tarea(&mut self.lista, id)
    }

    pub fn cargar(&mut self) {
        if let Ok(mut file) = fs::File::open(&self.archivo) {
            let mut data = String::new();
            if file.read_to_string(&mut data).is_ok() {
                if let Ok(lista_cargada) = serde_json::from_str(&data) {
                    self.lista = lista_cargada;
                } else {
                    println!("Error al leer {}. Se usará una lista vacía.", self.archivo);
                }
            }
        } else {
            println!("Archivo {} no encontrado. Se creará uno nuevo al guardar.", self.archivo);
        }
    }

    pub fn guardar(&self) {
        match serde_json::to_string_pretty(&self.lista) {
            Ok(data_json) => {
                match fs::File::create(&self.archivo) {
                    Ok(mut file) => {
                        if file.write_all(data_json.as_bytes()).is_ok() {
                            println!("Tareas guardadas exitosamente en {}.", self.archivo);
                        } else {
                            println!("Error al escribir en {}.", self.archivo);
                        }
                    },
                    Err(_) => println!("Error al crear el archivo {}.", self.archivo),
                }
            },
            Err(_) => println!("Error al serializar las tareas a JSON."),
        }
    }
}