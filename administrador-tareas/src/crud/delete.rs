
use crate::lista::Lista;

pub fn elimminar_tarea(lista: &mut Lista, id: String) -> bool {
    if let Some(index) = lista.contenido.iter().position(|t| t.id == id) {
        let tarea_eliminada = lista.contenido.remove(index);
        println!("Tarea eliminada: {}", tarea_eliminada.nombre);
        return true; 
    } else {
        println!("No se encontró la tarea con id: {}", id);
        return false; 
    }
}
#[cfg(test)]
mod tests {
    use super::elimminar_tarea;
    use crate::lista::Lista;
    use crate::tarea::{Tarea, EstadoTarea}; 
    
    fn crear_tarea_mock(id: &str, nombre: &str) -> Tarea {
        Tarea::new(
            id.to_string(), 
            nombre.to_string(), 
            0.0, 
            "2025-10-25".to_string() ,
        )
    }

    #[test]
    fn elimina_tarea_existente_correctamente() {
        let tarea_a = crear_tarea_mock("id-A", "Tarea A");
        let tarea_b = crear_tarea_mock("id-B", "Tarea B");
        let id_a_eliminar = "id-A".to_string();

        let mut lista = Lista {
            contenido: vec![tarea_a, tarea_b],
        };
        let longitud_inicial = lista.contenido.len();

        let resultado = elimminar_tarea(&mut lista, id_a_eliminar);

        assert!(resultado);
        assert_eq!(lista.contenido.len(), longitud_inicial - 1);
        assert_eq!(lista.contenido[0].id, "id-B");
    }

    #[test]
    fn no_elimina_tarea_inexistente() {
        let tarea_a = crear_tarea_mock("id-A", "Tarea A");
        let tarea_b = crear_tarea_mock("id-B", "Tarea B");
        let id_inexistente = "id-Z".to_string();

        let mut lista = Lista {
            contenido: vec![tarea_a.clone(), tarea_b.clone()],
        };
        let longitud_inicial = lista.contenido.len();

        let resultado = elimminar_tarea(&mut lista, id_inexistente);

        assert!(!resultado);
        assert_eq!(lista.contenido.len(), longitud_inicial);
        assert!(lista.contenido.iter().any(|t| t.id == "id-A"));
        assert!(lista.contenido.iter().any(|t| t.id == "id-B"));
    }

    #[test]
    fn no_elimina_de_lista_vacia() {
        let mut lista = Lista {
            contenido: Vec::new(),
        };
        let id_cualquiera = "cualquier-id".to_string();

        let resultado = elimminar_tarea(&mut lista, id_cualquiera);

        assert!(!resultado);
        assert!(lista.contenido.is_empty());
    }

    #[test]
    fn elimina_tarea_en_posicion_intermedia() {
        let tarea_a = crear_tarea_mock("id-A", "T A");
        let tarea_b = crear_tarea_mock("id-B", "T B");
        let tarea_c = crear_tarea_mock("id-C", "T C");

        let mut lista = Lista {
            contenido: vec![tarea_a, tarea_b, tarea_c],
        };

        let resultado = elimminar_tarea(&mut lista, "id-B".to_string());

        assert!(resultado);
        assert_eq!(lista.contenido.len(), 2);
        assert_eq!(lista.contenido[0].id, "id-A");
        assert_eq!(lista.contenido[1].id, "id-C");
    }
}