from flask import Flask, request

app = Flask(__name__)


@app.route('/')
def root():
    return "Hello from Scorer-Backend!"


if __name__ == '__main__':
    app.run(debug=True, port=9003)
