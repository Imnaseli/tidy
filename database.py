import sqlite3
from typing import List
import datetime
from model import Todo


con = sqlite3.connect("tidy.db")
c = con.cursor()

def createtable():
    c.execute("""CREATE TABLE IF NOT EXISTS todos(
                task text,
                category text,
                date_added text,
                date_completed text,
                status integer,
                position integer)""")

createtable()

def inserttodo(todo:Todo):
    c.execute('select count(*) FROM todos')
    count = c.fetchone()[0]
    todo.position = count if count else 0
    with con:
        c.execute('INSERT INTO todos VALUES(:task, :category, :date_added, :date_completed, :status, :position)',
        {'task':todo.task , 'category':todo.category, 'date_added':todo.date_added, 'date_completed':todo.date_completed, 'status':todo.status , 'position':todo.position})

def gettodo():
    c.execute('select * from todos')
    results = c.fetchall()
    todos = []
    for result in results:
        todos.append(Todo(*result))
    return todos

def deletetodo(position):
    c.execute("select count(*) from todos")
    count = c.fetchone()[0]
    with con:
        c.execute('DELETE FROM todos WHERE position =:position', {'position':position})
        for pos in range(position+1 , count):
            changeposition(pos, pos-1 ,False)

def changeposition(oldpos:int , newpos:int , commit=True):
    c.execute("UPDATE todos SET position =:newpos WHERE position=:oldpos",
    {'newpos':newpos , 'oldpos':oldpos})

    if commit:
        con.commit()

def updatetodo(pos:int , task:str , category:str):
    with con:
        if task is not None and category is not None:
            c.execute('UPDATE todos SET task=:task, category=:category WHERE position=:position',
            {'task':task , 'category':category, 'position':pos})

        elif task is not None:
            c.execute('UPDATE todos SET task=:task WHERE position=:position',
            {'task':task , 'position':pos})

        elif category is not None:
            c.execute('UPDATE todos SET category=:category WHERE position=:position',
            {'category':category, 'position':pos})

def completetodo(position:int):
    c.execute('UPDATE todos SET status = 2, date_completed=:date_completed WHERE position=:position' , 
    {"position":position , "date_completed":datetime.datetime.now().isoformat()})













