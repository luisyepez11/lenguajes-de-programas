
mod tarea;
mod lista;
mod interfaz;
mod crud; 

use lista::Lista;
use tarea::Tarea;
use interfaz::Interfaz;
use clap::{Parser, Subcommand};
use chrono::{Local, DateTime};

#[derive(Parser)]
#[command(author, version, about, long_about = None)]
#[command(propagate_version = true)]
struct Cli {
    #[command(subcommand)]
    command: Commands,
}

#[derive(Subcommand)]
enum Commands {
    Create {
        #[arg(short, long)]
        nombre: String,
        #[arg(short, long, default_value_t = 0.0)]
        porcentaje: f32,
        #[arg(short, long,default_value_t = get_current_date())]
        creacion: String,
    },
    List,
    ListComplit,
    ListNotInit,
    Update {
        #[arg(short, long)]
        id: String,
        #[arg(short = 'n', long)]
        nombre: Option<String>,
        #[arg(short = 'p', long)]
        porcentaje: Option<f32>,
    },
    Delete {
        #[arg(short, long)]
        id: String,
    },
    Filled {
        #[arg(short, long)]
        id: String,
    }
}

fn main() {
    let tareas: Vec<Tarea> = Vec::new();
    let lista: Lista = Lista::new(tareas);
    let mut interfaz = Interfaz::new("tarea.json".to_string(), lista);

    interfaz.cargar();

    let cli = Cli::parse();

    let mut cambios_realizados = false; 

    match &cli.command {
        Commands::Create { nombre, porcentaje, creacion } => {
            println!("--- Creando Tarea ---");
            interfaz.crear_tarea(
                nombre.clone(), 
                *porcentaje, 
                creacion.clone()
            );
            cambios_realizados = true;
        }
        
        Commands::List => {
            println!("--- Lista de Tareas ---");
            interfaz.ver_tareas();
        }

        Commands::ListComplit => {
            println!("--- Lista de Tareas Completadas ---");
            interfaz.ver_tareas_completadas();
        }

        Commands::ListNotInit => {
            println!("--- Lista de Tareas en Proceso---");
            interfaz.ver_tareas_no_iniciadas();
        }
        
        Commands::Update { id, nombre, porcentaje } => {
            println!("--- Actualizando Tarea {} ---", id);
            if let Some(n) = nombre {
                if interfaz.actualizar_tarea_nombre(id.clone(), n.clone()) {
                    cambios_realizados = true;
                }
            }
            if let Some(p) = porcentaje {
                if interfaz.actualizar_tarea_porcentaje(id.clone(), *p) {
                    cambios_realizados = true;
                }
            }
            if nombre.is_none() && porcentaje.is_none() {
                println!("No especificaste qué actualizar. Usa --nombre ('-n') o --porcentaje ('-p').");
            }
        }
        
        Commands::Delete { id } => {
            println!("--- Eliminando Tarea {} ---", id);
            if interfaz.eliminar_tarea(id.clone()) {
                cambios_realizados = true;
            }
        }

        Commands::Filled { id} =>{
            println!("--- Completando Tarea {} ---", id);
            if interfaz.actualizar_tarea_porcentaje_completa(id.clone()){
                cambios_realizados =true;
            }
        }
    }

    if cambios_realizados {
        println!("\n--- Guardando cambios en {} ---", interfaz.archivo);
        interfaz.guardar();
    } else {
        println!("\nNo se realizaron cambios.");
    }
}
fn get_current_date() -> String {
    let now: DateTime<Local> = Local::now();
    now.format("%Y-%m-%d %H:%M:%S").to_string()
}