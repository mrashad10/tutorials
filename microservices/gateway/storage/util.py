import pika
import json

def upload(f, fs, channel, access):
    """
    Uploads a file to GridFS and sends a message to RabbitMQ for further processing.

    Args:
        f: The file to be uploaded.
        fs: GridFS instance for storing the file.
        channel: RabbitMQ channel for message publishing.
        access: User access information.

    Returns:
        tuple: (success_message, error_message) where success_message is a success message if
               upload and message sending are successful, and error_message is an error message
               and status code if any step fails.
    """
    try:
        fid = fs.put(f)
    except Exception as err:
        print(err)
        return "internal server error", 500

    message = {
        "video_fid": str(fid),
        "mp3_fid": None,
        "username": access["username"],
    }

    try:
        channel.basic_publish(
            exchange="",
            routing_key="video",
            body=json.dumps(message),
            properties=pika.BasicProperties(
                delivery_mode=pika.spec.PERSISTENT_DELIVERY_MODE
            ),
        )
    except Exception as err:
        print(err)
        fs.delete(fid)
        return "internal server error", 500

    return "success", None
