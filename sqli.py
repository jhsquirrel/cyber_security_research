from flask import Flask, request, render_template_string, render_template
import jinja2
import os
import sqlite3

app = Flask(__name__)

DB1 = "./my.db"

# start with python3 sqli.py

def getData(user):
    try:
        conn = sqlite3.connect(DB1)
        c = conn.cursor()
        try:
            c.execute("SELECT * FROM users WHERE name = ?", (user,))
            r = c.fetchall()
            conn.close()
            print("getuser r", r)
            return r
        except Exception as e:
            print(str(e))
            conn.close()
    except Exception as e:
        print(str(e))
    return None 

def getDataBad(user):
    try:
        conn = sqlite3.connect(DB1)
        c = conn.cursor()
        print("getuserbad=",user)
        try:
            s = "SELECT * FROM users WHERE name = '" + user + "';"
            print("getuserbad s", s)
            c.execute(s)
            r = c.fetchall()
            conn.close()
            print("getuserbad r", r)
            return r
        except Exception as e:
            print(str(e))
            conn.close()
    except Exception as e:
        print(str(e))
    return None

@app.route("/")
def login():
    return render_template("sqli.html")

@app.route("/userok")
def user():
    name = request.values.get('name')
    r = getData(name)
    print("user", name)
    template = '''%s''' % r
    return render_template_string(template, r=r)

# used to be called userbad
@app.route("/user")
def userbad():
    name = request.values.get('name')
    r = getDataBad(name)
    print("userbad", name)
    template = '''%s''' % r
    print("userbad", template)
    return render_template_string(template, r=r)

if __name__=="__main__":
    app.run(debug=True)
