from flask import Flask, jsonify

app = Flask(__name__)

health_status = True

@app.route('/hz')
def health():
    return "OK", 200

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=8080)