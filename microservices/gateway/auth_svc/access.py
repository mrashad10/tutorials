import os
import requests

def login(request):
    """
    Logs in a user by sending a request to the authentication service.

    Args:
        request: Flask request object containing authorization information.

    Returns:
        tuple: (token, error) where token is the JWT token received upon successful login and
               error is a tuple containing an error message and status code if login fails,
               otherwise None.
    """
    # Get authorization from request
    auth = request.authorization

    # Check if authorization exists
    if not auth:
        return None, ("missing credentials", 401)

    # Create basic authentication tuple
    basicAuth = (auth.username, auth.password)

    # Send login request to authentication service
    response = requests.post(
        f"http://{os.environ.get('AUTH_SVC_ADDRESS')}/login", auth=basicAuth
    )

    # Check if login was successful
    if response.status_code == 200:
        return response.text, None
    else:
        return None, (response.text, response.status_code)
