import jwt
import datetime
import os
from flask import Flask, request
from flask_mysqldb import MySQL
from werkzeug.security import check_password_hash, generate_password_hash

server = Flask(__name__)
mysql = MySQL(server)

# Config (with default values)
server.config["MYSQL_HOST"] = os.environ.get("MYSQL_HOST", "localhost")
server.config["MYSQL_USER"] = os.environ.get("MYSQL_USER", "root")
server.config["MYSQL_PASSWORD"] = os.environ.get("MYSQL_PASSWORD", "password")
server.config["MYSQL_DB"] = os.environ.get("MYSQL_DB", "my_database")
server.config["MYSQL_PORT"] = int(os.environ.get("MYSQL_PORT", 3306))
server.config["JWT_SECRET"] = os.environ.get("JWT_SECRET", "my_secret_key")


def get_db():
    """Helper function to get a database cursor."""
    return mysql.connection.cursor()


@server.route("/login", methods=["POST"])
def login():
    """Endpoint for user login."""
    auth = request.authorization
    if not auth:
        return "Missing credentials", 401

    try:
        with get_db() as cur:
            res = cur.execute(
                "SELECT email, password FROM user WHERE email=%s", (auth.username,)
            )

            if res > 0:
                user_row = cur.fetchone()
                email = user_row[0]
                password = user_row[1]

                if auth.username != email or not check_password_hash(password, auth.password):
                    return "Invalid credentials", 401
                else:
                    return create_jwt(auth.username, server.config["JWT_SECRET"], True)
    except Exception as e:
        return str(e), 500

    return "Invalid credentials", 401


@server.route("/validate", methods=["POST"])
def validate():
    """Endpoint for validating JWT."""
    encoded_jwt = request.headers.get("Authorization")

    if not encoded_jwt:
        return "Missing credentials", 401

    encoded_jwt = encoded_jwt.split(" ")[1]

    try:
        decoded = jwt.decode(
            encoded_jwt, server.config["JWT_SECRET"], algorithms=["HS256"]
        )
        return decoded, 200
    except jwt.ExpiredSignatureError:
        return "Token has expired", 401
    except jwt.DecodeError:
        return "Invalid token", 401

    return "Not authorized", 403


def create_jwt(username, secret, authz):
    """Create a JWT token."""
    return jwt.encode(
        {
            "username": username,
            "exp": datetime.datetime.now(tz=datetime.timezone.utc)
            + datetime.timedelta(days=1),
            "iat": datetime.datetime.utcnow(),
            "admin": authz,
        },
        secret,
        algorithm="HS256",
    )


if __name__ == "__main__":
    server.run(host="0.0.0.0", port=5000)
