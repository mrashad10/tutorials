import os
import requests

def token(request):
    """
    Validates a JWT token by sending a request to the authentication service.

    Args:
        request: Flask request object containing headers.

    Returns:
        tuple: (token, error) where token is the validated JWT token and error is a tuple
               containing an error message and status code if validation fails, otherwise None.
    """
    if not "Authorization" in request.headers:
        return None, ("missing credentials", 401)

    token = request.headers["Authorization"]

    if not token:
        return None, ("missing credentials", 401)

    response = requests.post(
        f"http://{os.environ.get('AUTH_SVC_ADDRESS')}/validate",
        headers={"Authorization": token},
    )

    if response.status_code == 200:
        return response.text, None
    else:
        return None, (response.text, response.status_code)
