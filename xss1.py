from flask import Flask, request, render_template_string, render_template

app = Flask(__name__)

@app.route("/")
def login():
    return render_template("xss1.html")

@app.route("/page")
def page():
    name = ""
    password = ""
    if request.args.get('name'):
        print(request.args.get('name'))
        print(request.args.get('password'))
        name = request.args.get('name')
        password = request.args.get('password')
    template = '''hello %s''' % name
    return render_template_string(template, name=name)

if __name__=="__main__":
    app.run(debug=True)
