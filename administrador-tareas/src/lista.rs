use crate::tarea::Tarea; 
use serde::{Serialize, Deserialize};
use std::fs;

#[derive(Serialize, Deserialize, Debug)]
pub struct Lista {
    pub contenido :  Vec<Tarea>
}

impl Lista {
    pub fn new(contenido: Vec<Tarea>) -> Lista{
        Lista {
            contenido
        }
    }

    pub fn escribir(&self, archivo: String) -> Result<(), Box<dyn std::error::Error>> {
        let json_data = serde_json::to_string_pretty(&self.contenido)?;
        fs::write(archivo, json_data)?;
        Ok(())
    }

    pub fn cargar_datos(&mut self, archivo: String) -> Result<(), Box<dyn std::error::Error>> {
        let json_string = match fs::read_to_string(archivo) {
            Ok(s) => s,
            Err(e) if e.kind() == std::io::ErrorKind::NotFound => {
                println!("Archivo no encontrado, se creará uno nuevo al guardar.");
                return Ok(());
            },
            Err(e) => return Err(Box::new(e)),
        };

        if json_string.trim().is_empty() {
            println!("Archivo vacío, iniciando con lista vacía.");
            self.contenido = Vec::new();
            return Ok(());
        }

        self.contenido = serde_json::from_str(&json_string)?;
        println!("Tareas cargadas exitosamente");
        Ok(())
    }
}