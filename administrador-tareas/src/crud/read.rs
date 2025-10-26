use crate::lista::Lista;

pub fn ver_tareas(lista: &Lista) {
    if lista.contenido.is_empty() {
        println!("No hay tareas para mostrar.");
        return;
    }
    for tarea in &lista.contenido {
        tarea.mostrar_tarea();
    }
}

pub fn ver_tareas_completadas(lista: &Lista) {
    if lista.contenido.is_empty() {
        println!("No hay tareas para mostrar.");
        return;
    }
    for tarea in &lista.contenido {
        if tarea.porcentaje == 100.00{
            tarea.mostrar_tarea();
        }
    }
}
pub fn ver_tareas_no_iniciadas(lista: &Lista) {
    if lista.contenido.is_empty() {
        println!("No hay tareas para mostrar.");
        return;
    }
    for tarea in &lista.contenido {
        if tarea.porcentaje < 1.00{
            tarea.mostrar_tarea();
        }
    }
}
#[cfg(test)]
mod tests {
    use super::ver_tareas;
    use crate::lista::Lista;
    use crate::tarea::{Tarea}; 
    fn crear_tarea_mock(id: &str, nombre: &str) -> Tarea {
        Tarea::new(
            id.to_string(), 
            nombre.to_string(), 
            50.0, 
            "2025-10-25".to_string() 
        )
    }

    #[test]
    fn ver_tareas_imprime_mensaje_si_lista_esta_vacia() {
        let lista = Lista {
            contenido: Vec::new(),
        };


        ver_tareas(&lista);
    }
    
    #[test]
    fn ver_tareas_imprime_todas_las_tareas() {
        let tarea1 = crear_tarea_mock("id-1", "Tarea Uno");
        let tarea2 = crear_tarea_mock("id-2", "Tarea Dos");

        let lista = Lista {
            contenido: vec![tarea1, tarea2],
        };
        ver_tareas(&lista);
    }

    #[test]
    fn ver_tareas_imprime_una_sola_tarea() {
        let tarea = crear_tarea_mock("id-u", "Tarea Unica");

        let lista = Lista {
            contenido: vec![tarea],
        };

        ver_tareas(&lista);
    }
}