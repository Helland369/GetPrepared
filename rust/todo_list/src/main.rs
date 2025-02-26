use std::io;

struct Task {
    id: u32,
    description: String,
    compleated: bool,
}

struct TaskList {
    tasks: Vec<Task>,
}

impl TaskList {
    fn new() -> Self {
        TaskList { tasks: Vec::new() }
    }

    fn add_task(&mut self, description: &str) {
        let id = (self.tasks.len() as u32) + 1;
        let _new_taks = Task {
            id: id,
            description: description.to_string(),
            compleated: false,
        };
        self.tasks.push(_new_taks);
    }

    fn list_tasks(&self) {
        for task in &self.tasks {
            println!(
                "Id: {}, Description: {}, Compleated: {}",
                task.id, task.description, task.compleated
            );
        }
    }

    fn complete_task(&mut self, id: u32) {
        if let Some(task) = self.tasks.iter_mut().find(|x| x.id == id) {
            task.compleated = true;
        }
    }

    fn remove_taks(&mut self, id: u32) {
        self.tasks.retain(|task| task.id != id);
    }
}

fn main() {
    let mut task_list = TaskList::new();

    let select: i8 = -1;
    let mut input = String::new();

    while select != 0 {
        println!(
            "[1] Add new task\n[2] List all tasks\n[3] Complete a task\n[4]Remove task\n[0] Exit"
        );
        
        input.clear();
        io::stdin()
            .read_line(&mut input)
            .expect("Failed to read line!");
        let select: i8 = match input.trim().parse() {
            Ok(num) => num,
            Err(_) => continue,
        };

        match select {
            0 => break,
            1 => {
                println!("Write your task: ");
                input.clear();
                io::stdin()
                    .read_line(&mut input)
                    .expect("Failed to read line!");
                task_list.add_task(&input.trim());
            }
            2 => task_list.list_tasks(),
            3 => {
                println!("Enter the task ID to complete: ");
                input.clear();
                io::stdin()
                    .read_line(&mut input)
                    .expect("Failed to read line!");
                let id: u32 = match input.trim().parse() {
                    Ok(num) => num,
                    Err(_) => {
                        println!("Invalid ID! Please enter a valid number.");
                        continue;
                    }
                };
                task_list.complete_task(id);
            }
            4 => {
                println!("Enter the task ID to remove: ");
                input.clear();
                io::stdin()
                    .read_line(&mut input)
                    .expect("Failed to read line!");
                let id: u32 = match input.trim().parse() {
                    Ok(num) => num,
                    Err(_) => {
                        println!("Invalid ID! Please enter a valid number.");
                        continue;
                    }
                };
                task_list.remove_taks(id);
            }
            _ => println!("Invalid option! Please choose a valid number."),
        }
    }
}
