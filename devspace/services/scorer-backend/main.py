from flask import Flask, request

app = Flask(__name__)


@app.route('/')
def hello_world():
    return "Hello From Devspace Scorer-Backend!"


if __name__ == '__main__':
    app.run(debug=True, port=9003)
