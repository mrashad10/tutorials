import pika
import sys
import os
import time
from send import email


def main():
    """
    Main function to process MP3 conversion tasks and send notifications.

    Returns:
        None
    """
    # RabbitMQ connection
    connection = pika.BlockingConnection(pika.ConnectionParameters(host="rabbitmq"))
    channel = connection.channel()

    def callback(ch, method, properties, body):
        """
        Callback function to handle incoming messages from RabbitMQ.

        Args:
            ch: Channel object.
            method: Delivery method.
            properties: Message properties.
            body: Message body (JSON message).
        Returns:
            None
        """
        err = email.notification(body)
        if err:
            ch.basic_nack(delivery_tag=method.delivery_tag)
        else:
            ch.basic_ack(delivery_tag=method.delivery_tag)

    # Consume messages from the specified queue
    channel.basic_consume(
        queue=os.environ.get("MP3_QUEUE"), on_message_callback=callback
    )

    print("Waiting for messages. To exit press CTRL+C")

    # Start consuming messages
    channel.start_consuming()


if __name__ == "__main__":
    try:
        main()
    except KeyboardInterrupt:
        print("Interrupted")
        try:
            sys.exit(0)
        except SystemExit:
            os._exit(0)
