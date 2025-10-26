use crate::lista::Lista;

pub fn actualizar_tarea_nombre(lista: &mut Lista, id: String, nombre: String) -> bool {
    for tarea in &mut lista.contenido {
        if tarea.id == id {
            tarea.cambiar_nombre(nombre.clone());
            println!("Nombre de tarea actualizado.");
            return true;
        }
    }
    println!("No se encontró la tarea con id: {}", id);
    false 
}

pub fn actualizar_tarea_porcentaje(lista: &mut Lista, id: String, porcentaje: f32) -> bool {
    for tarea in &mut lista.contenido {
        if tarea.id == id {
            tarea.cambiar_porcentaje(porcentaje);
            println!("Porcentaje de tarea actualizado.");
            return true; 
        }
    }
    println!("No se encontró la tarea con id: {}", id);
    false 
}

pub fn actualizar_tarea_porcentaje_completa(lista: &mut Lista, id: String) -> bool {
    for tarea in &mut lista.contenido {
        if tarea.id == id {
            tarea.cambiar_porcentaje(100.00);
            println!("Porcentaje de tarea actualizado.");
            return true; 
        }
    }
    println!("No se encontró la tarea con id: {}", id);
    false 
}
#[cfg(test)]
mod tests {
    use super::{actualizar_tarea_nombre, actualizar_tarea_porcentaje, actualizar_tarea_porcentaje_completa};
    use crate::lista::Lista;
    use crate::tarea::{Tarea, EstadoTarea};

    fn crear_tarea_mock(id: &str, nombre: &str, porcentaje: f32) -> Tarea {
        Tarea::new(
            id.to_string(), 
            nombre.to_string(), 
            porcentaje, 
            "2025-01-01".to_string()
        )
    }

    // -------------------------------------------------------------------------
    // Pruebas para actualizar_tarea_nombre
    // -------------------------------------------------------------------------

    #[test]
    fn actualizar_nombre_existe_retorna_true_y_cambia() {
        let mut lista = Lista {
            contenido: vec![crear_tarea_mock("id1", "Original", 10.0)],
        };
        let nuevo_nombre = "Nuevo Nombre".to_string();

        let resultado = actualizar_tarea_nombre(&mut lista, "id1".to_string(), nuevo_nombre.clone());

        assert!(resultado);
        assert_eq!(lista.contenido[0].nombre, nuevo_nombre);
    }

    #[test]
    fn actualizar_nombre_no_existe_retorna_false() {
        let mut lista = Lista {
            contenido: vec![crear_tarea_mock("id1", "Original", 10.0)],
        };

        let resultado = actualizar_tarea_nombre(&mut lista, "id_x".to_string(), "Nuevo".to_string());

        assert!(!resultado);
        assert_eq!(lista.contenido[0].nombre, "Original"); // Verifica que no se modificó
    }
    
    // -------------------------------------------------------------------------
    // Pruebas para actualizar_tarea_porcentaje
    // -------------------------------------------------------------------------

    #[test]
    fn actualizar_porcentaje_existe_retorna_true_y_cambia() {
        let mut lista = Lista {
            contenido: vec![crear_tarea_mock("id2", "Tarea", 10.0)],
        };
        let nuevo_porcentaje = 55.5;

        let resultado = actualizar_tarea_porcentaje(&mut lista, "id2".to_string(), nuevo_porcentaje);

        assert!(resultado);
        assert_eq!(lista.contenido[0].porcentaje, nuevo_porcentaje);
        // Verifica que el estado también se actualizó (asumiendo lógica de Tarea)
        assert_eq!(lista.contenido[0].estado, EstadoTarea::Avanzado); 
    }

    #[test]
    fn actualizar_porcentaje_no_existe_retorna_false() {
        let mut lista = Lista {
            contenido: vec![crear_tarea_mock("id2", "Tarea", 10.0)],
        };

        let resultado = actualizar_tarea_porcentaje(&mut lista, "id_x".to_string(), 99.0);

        assert!(!resultado);
        assert_eq!(lista.contenido[0].porcentaje, 10.0); // Verifica que no se modificó
    }


    #[test]
    fn actualizar_completa_existe_retorna_true_y_es_cien() {
        let mut lista = Lista {
            contenido: vec![crear_tarea_mock("id3", "Tarea", 10.0)],
        };

        let resultado = actualizar_tarea_porcentaje_completa(&mut lista, "id3".to_string());

        assert!(resultado);
        assert_eq!(lista.contenido[0].porcentaje, 100.00);
        assert_eq!(lista.contenido[0].estado, EstadoTarea::Completado);
    }

    #[test]
    fn actualizar_completa_no_existe_retorna_false() {
        let mut lista = Lista {
            contenido: vec![crear_tarea_mock("id3", "Tarea", 10.0)],
        };

        let resultado = actualizar_tarea_porcentaje_completa(&mut lista, "id_x".to_string());

        assert!(!resultado);
        assert_eq!(lista.contenido[0].porcentaje, 10.0); // Verifica que no se modificó
    }
}