from flask import Flask, redirect, render_template, request
from httpx import get, post

app = Flask(__name__)


@app.route("/")
def index():
    return render_template("index.html")


@app.route("/course", methods=["GET"])
def course():
    data = get("http://localhost:8080/course")
    return render_template("course.html", data=data)


@app.route("/register", methods=["GET", "POST"])
def register(request):
    if request.method == "POST":
        first_name = request.form.get("first_name")
        last_name = request.form.get("last_name")
        password = request.form.get("password")
        email = request.form.get("email")
        phone = request.form.get("phone")
        address = request.form.get("address")
        city = request.form.get("city")
        postcode = request.form.get("postcode")
        country_code = request.form.get("country_code")
        json_data = {
            "first_name": first_name,
            "last_name": last_name,
            "password": password,
            "email": email,
            "phone": phone,
            "address": address,
            "city": city,
            "postcode": postcode,
            "country_code": country_code,
        }
        try:
            res = post("http://localhost:8080/register", json=json_data)
            return redirect("/dashboard")
        except:
            return render_template("error.html")
    return render_template("register.html")


@app.route("/login", methods=["GET", "POST"])
def user_login():
    if request.method == "POST":
        username = request.form.get("username")
        password = request.form.get("password")
        json_data = {"username": username, "password": password}
        res = post("http://localhost:8080/login")
    return render_template("login.html")


@app.route("/dashboard", methods=["GET"])
def dashboard():
    # harus ambil data account, hit apinya
    return render_template("dashboard.html")


@app.route("/course/:id", methods=["GET"])
def watch_course():
    return
