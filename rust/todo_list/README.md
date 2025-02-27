# Simple todo app in rust

- [Run and install](#How-to-run-and-install-the-program)
- [About](#About-the-program)
    - [Struct](#Struct)
    - [Vector](#Vector)
    - [Add task](#Add-task)
    - [Print all tasks](#print-all-tasks)
    - [Complete task](#Complete-task)
    - [Remove task](#Remove-task)

# How to run and install the program


First clone the repository:

```bash

git clone https://github.com/Helland369/GetPrepared.git

```

Then cd in to the todo_list directory:

```bash

cd GetPrepared/rust/todo_list

```

If you have Rust installed, you can simply run the program. If not, downloade Rust [here](https://www.rust-lang.org/). You can also intall Rust with your favorite package manager (pacman, apt, etc..).

```bash

cargo run

```

# About the program


### Struct

There is a struct "Task", which represents the properties of a task.

```rust

struct Task {
    id: u32,
    description: String,
    completed: bool,
}

```

### Vector

Each task is stored in a vector inside a struct, making it a list or array of tasks.

```rust

struct TaskList {
    tasks: Vec<Task>,
}

```

### Add task

The add_task function adds a new task to the vector, the function takes one argument in form of a new task description (string).

```rust

fn add_task(&mut self, description: &str) {
    let id = (self.tasks.len() as u32) + 1;
    let _new_task = Task {
        id: id,
        description: description.to_string(),
        compleated: false,
    };
    self.tasks.push(_new_taks);
}

```

### Print all tasks

The list_tasks function takes no arguments. It's job is to print/display all the tasks in the vector.

```rust


fn list_tasks(&self) {
    for task in &self.tasks {
        println!(
            "Id: {}, Description: {}, Compleated: {}",
            task.id, task.description, task.compleated
        );
    }
}

```

### Complete task

The complete_task function takes an id as an argument and updates the corresponding task's boolean value from false to true.

```rust

fn complete_task(&mut self, id: u32) {
    if let Some(task) = self.tasks.iter_mut().find(|x| x.id == id) {
        task.compleated = true;
    }
}

```

### Remove task

The remove_taks function remove a task from the vector of tasks.

```rust

fn remove_taks(&mut self, id: u32) {
    self.tasks.retain(|task| task.id != id);
}

```

