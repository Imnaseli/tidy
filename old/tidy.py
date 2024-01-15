from click import pass_context
import typer
from model import Todo
from database import *
from rich.console import Console
from rich.table import Table

console = Console()
app = typer.Typer()

@app.command(short_help="Add a Task")
def add ():
    task = input("Add title of task: ")
    category = input('Add title of task category: ')
    typer.echo(f"adding {task} to the {category} category")
    todo = Todo(task , category  ) 
    inserttodo(todo)
    show()

@app.command(short_help="Delete a Task")
def delete ():
    show()
    position = int(input("What is the position of the task you want to delete: "))
    
    #TODO: stopped here, Change the echo to tthe title task
    todos = gettodo()
    todos[position]
    typer.echo(f"deleting task {position}")
    deletetodo(position - 1) 
    show()


#TODO: format the whole update command 
@app.command(short_help="Update a Task")
def update ():
    show()
    position = int(input("What task do you want to update (Use the position): ")) - 1
    option = int(input("""what do you want to update,
1. category
2. tasks
                :"""))

    if option == 1:
        newcategory = input("Enter new updated category: ")
        updatecategory(position , newcategory)
    elif option == 2:
        newtask = input("Enter new updated task: ")
        updatetask(position , newtask)
    else:
        print("Please enter one of the available options")
        return

    typer.echo(f"Updating task {position}")
    show()


@app.command(short_help="Complete a Task")
def complete ():
    position = int(input("What task did you complete (Use the position): ")) - 1
    completetodo(position)
    show()

@app.command(short_help="Show all Tasks")
def show():
    tasks = gettodo() 
    console.print("[bold magenta]Tasks[/bold magenta]","" , justify="center")

    table = Table(show_header=True, header_style="bold magenta" ,) 
    table.add_column("Position" , style="dim" , width=10)
    table.add_column("Task" , min_width=20)
    table.add_column("Category" , min_width=14 , justify="right")
    table.add_column("Status" , min_width=12 , justify="right")
    
    table.add_column("Date Completed" , min_width=15 , justify="right")
        
    def categorycolor (color):
        colors = {"learn":'cyan' , 'youtube':'red' , 'sports':'cyan' , 'food':'green'}
        color = color.lower()
        if color in colors:
            return colors[color]
        return "white"


    for idx , task in enumerate(tasks , start=1):
        c = categorycolor(task.category)
        isdonestr = "Done" if task.status == 2 else "Not Done"
        #isdonestr = str(task.status)       
        table.add_row(str(idx) , task.task , f"[{c}]{task.category}[/{c}]" , isdonestr , task.date_completed)

    console.print(table)   

if __name__ == "__main__":
    app()


