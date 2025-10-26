use serde::{Serialize, Deserialize};

#[derive(Serialize, Deserialize, Clone, Debug, PartialEq)]
pub enum EstadoTarea {
    Incompleto,
    EnProceso,
    Avanzado,
    Completado,
}

#[derive(Serialize, Deserialize, Clone, Debug)]
pub struct Tarea {
    pub id: String,
    pub nombre: String,
    pub porcentaje: f32,
    pub estado: EstadoTarea,
    pub creacion: String, 
}

impl Tarea {
    pub fn new(id: String,nombre: String,porcentaje: f32,creacion: String) -> Tarea{
        Tarea{
            id,
            nombre,
            porcentaje,
            creacion,
            estado:EstadoTarea::Incompleto
        }
    }

    pub fn mostrar_tarea(&self){
        println!("-----------------------------------------------");
        println!("Nombre de la tarea : {}",self.nombre);
        println!("Su id es : {}",self.id);
        println!("Porcentaje completado:");
        let num = self.porcentaje * 2.0; 
        let mut numeral = "#".to_string();
        let mut sumador = "[".to_string();
        for _ in 0..(num/10.0) as i32 {
            sumador += &numeral;
        }
        numeral = "_".to_string();
        if num < 200.0{
            for _ in 0..(20.0-(num/10.0)) as i32 {
                sumador += &numeral;
            }
        }
        sumador += "]";
        println!("{} {:.2}%",sumador,self.porcentaje);
        println!("{:?}",self.estado);
        println!("-----------------------------------------------");

    }

    pub fn cambiar_nombre(&mut self,nuevo_nombre: String){
        self.nombre=nuevo_nombre;
        self.mostrar_tarea();
    }

    pub fn cambiar_porcentaje(&mut self,nuevo_porcentaje: f32){
        self.porcentaje = nuevo_porcentaje;
        self.actualizar_estado_por_porcentaje();
        self.mostrar_tarea();
    }
    pub fn actualizar_estado_por_porcentaje(&mut self) {
        match self.porcentaje {
            p if p == 0.0 => {
                self.estado = EstadoTarea::Incompleto;
            }
            p if p > 0.0 && p < 50.0 => {
                self.estado = EstadoTarea::EnProceso;
            }
            p if p >= 50.0 && p < 100.0 => {
                self.estado = EstadoTarea::Avanzado;
            }
            p if p == 100.0 => {
                self.estado = EstadoTarea::Completado;
            }
            _ => {
                self.estado = EstadoTarea::Incompleto;
            }
        }
    }
}