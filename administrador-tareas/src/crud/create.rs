use rand::RngCore;

use crate::lista::Lista;
use crate::tarea::Tarea;

pub fn crear_tarea(lista: &mut Lista, nombre: String, porcentaje: f32, creacion: String) {
    let mut bytes = vec![0u8; 8];
    rand::thread_rng().fill_bytes(&mut bytes);
    let id_hex = hex::encode(bytes);
    let mut tarea = Tarea::new(id_hex, nombre, porcentaje,creacion);
    tarea.actualizar_estado_por_porcentaje();
    lista.contenido.push(tarea.clone());
    println!("--- Tarea Creada ---");
    tarea.mostrar_tarea();
}
#[cfg(test)]
mod tests {
    use super::*; 
    use crate::lista::Lista;
    
    #[test]
    fn crea_y_agrega_tarea_correctamente() {
        let mut lista = Lista::new(Vec::new());
        let nombre_tarea = "Hacer ejercicio".to_string();
        let porcentaje_tarea = 75.00;
        let creacion_tarea = "2025-10-24".to_string();

        let longitud_inicial = lista.contenido.len();

        crear_tarea(&mut lista, nombre_tarea.clone(), porcentaje_tarea, creacion_tarea.clone());

        assert_eq!(lista.contenido.len(), longitud_inicial + 1);

        let tarea_anadida = lista.contenido.last().unwrap();

        assert_eq!(tarea_anadida.nombre, nombre_tarea);
        assert_eq!(tarea_anadida.porcentaje, porcentaje_tarea);
        assert_eq!(tarea_anadida.creacion, creacion_tarea);
        assert!(!tarea_anadida.id.is_empty());
        assert_eq!(tarea_anadida.id.len(), 16);
    }

    #[test]
    fn maneja_lista_vacia() {
        let mut lista = Lista {
            contenido: Vec::new(),
        };

        crear_tarea(&mut lista, "Tarea Inicial".to_string(), 0.0, "Ahora".to_string());
        
        assert_eq!(lista.contenido.len(), 1);
    }

    #[test]
    fn genera_ids_unicos_para_multiples_tareas() {
        let mut lista = Lista {
            contenido: Vec::new(),
        };

        crear_tarea(&mut lista, "T1".to_string(), 10.0, "D1".to_string());
        crear_tarea(&mut lista, "T2".to_string(), 90.00, "D2".to_string());
        
        assert_ne!(lista.contenido[0].id, lista.contenido[1].id);
    }
}
