from click import pass_context
import typer

from model import Todo
from database import *

from rich.console import Console
from rich.table import Table

console = Console()
app = typer.Typer()

@app.command(short_help="Add a Task")
def add (task:str , category:str):
    typer.echo(f"adding {task} to the {category} category")
    todo = Todo(task , category) 
    inserttodo(todo)
    show()

@app.command(short_help="Delete a Task")
def delete (position:int):
    typer.echo(f"deleting task {position}")
    deletetodo(position - 1) 
    show()


@app.command(short_help="Update a Task")
def update (position:int , task:str = None , category:str = None):
    typer.echo(f"Updating task {position}")
    updatetodo(position - 1 , task , category)
    show()


@app.command(short_help="Complete a Task")
def complete (position:int):
    typer.echo(f"deleting task {position}")
    completetodo(position - 1)
    show()

@app.command(short_help="Show all Tasks")
def show():
    #tasks = [("Todo1" , "study"),("Todo2" , "Dance")]
    tasks = gettodo() 
    console.print("[bold magenta]Tasks[/bold magenta]","" , justify="center")

    table = Table(show_header=True, header_style="bold magenta" ,) 
    table.add_column("#" , style="dim" , width=6)
    table.add_column("Task" , min_width=20)
    table.add_column("Category" , min_width=14 , justify="right")
    table.add_column("Status" , min_width=12 , justify="right")
        
    def categorycolor (color):
        colors = {"learn":'cyan' , 'youtube':'red' , 'sports':'cyan' , 'food':'green'}
        color = color.lower()
        if color in colors:
            return colors[color]
        return "white"


    for idx , task in enumerate(tasks , start=1):
        c = categorycolor(task.category)
        #isdonestr = "Done" if task.status == 2 else "Not Done"
        isdonestr = str(task.status)       
        table.add_row(str(idx) , task.task , f"[{c}]{task.category}[/{c}]" , isdonestr)

    console.print(table)   

if __name__ == "__main__":
    app()


