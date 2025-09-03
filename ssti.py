from flask import Flask, request, render_template_string, render_template
import jinja2
import os

app = Flask(__name__)

# start with python3 ssti.py

@app.route("/")
def login():
    return render_template("ssti.html")

@app.route("/page")
def page():
    name = request.values.get('name')
    print("PA", name)
    #output = jinja2.from_string('Hello ' + name + '!').render()
    template = "Hello " + name + "!"
    return render_template_string(template)

@app.route('/vuln')
def hello_ssti():
    person = {'name':"HackerTHM", 'password':"123456789"}
    print("A", request.args)
    if request.args.get('name'):
        print("B", request.args.get('name'))
        person['name'] = request.args.get('name')
        print("C", person['name'])

    template = '''<h2>Hello %s!</h2>''' % person['name'] # !LOL
    print("D", template)

    return render_template_string(template, person=person)

def get_user_file(f_name):
    with open(f_name) as f:
        return f.readlines()

app.jinja_env.globals['get_user_file'] = get_user_file

if __name__=="__main__":
    app.run(debug=True)

