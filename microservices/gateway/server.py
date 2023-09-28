import os
import gridfs
import pika
import json
from flask import Flask, request, send_file
from flask_pymongo import PyMongo
from auth import validate
from auth_svc import access
from storage import util
from bson.objectid import ObjectId

server = Flask(__name__)

# MongoDB configurations for videos and mp3s
mongo_video = PyMongo(server, uri="mongodb://host.minikube.internal:27017/videos")
mongo_mp3 = PyMongo(server, uri="mongodb://host.minikube.internal:27017/mp3s")

# GridFS for storing videos and mp3s
fs_videos = gridfs.GridFS(mongo_video.db)
fs_mp3s = gridfs.GridFS(mongo_mp3.db)

# RabbitMQ connection
connection = pika.BlockingConnection(pika.ConnectionParameters("rabbitmq"))
channel = connection.channel()

@server.route("/login", methods=["POST"])
def login():
    """
    Endpoint for user login.

    Returns:
        str: JWT token if login is successful.
    """
    token, err = access.login(request)

    if not err:
        return token
    else:
        return err

@server.route("/upload", methods=["POST"])
def upload():
    """
    Endpoint for uploading files.

    Returns:
        str: Success message if upload is successful.
    """
    access, err = validate.token(request)

    if err:
        return err

    access = json.loads(access)

    if access["admin"]:
        if len(request.files) > 1 or len(request.files) < 1:
            return "exactly 1 file required", 400

        for _, f in request.files.items():
            err = util.upload(f, fs_videos, channel, access)

            if err:
                return err

        return "success!", 200
    else:
        return "not authorized", 401

@server.route("/download", methods=["GET"])
def download():
    """
    Endpoint for downloading files.

    Returns:
        file: Downloadable file if user has admin access.
    """
    access, err = validate.token(request)

    if err:
        return err

    access = json.loads(access)

    if access["admin"]:
        fid_string = request.args.get("fid")

        if not fid_string:
            return "fid is required", 400

        try:
            out = fs_mp3s.get(ObjectId(fid_string))
            return send_file(out, download_name=f"{fid_string}.mp3")
        except Exception as err:
            print(err)
            return "internal server error", 500

    return "not authorized", 401

if __name__ == "__main__":
    server.run(host="0.0.0.0", port=8080)
