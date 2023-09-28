# Use Python 3.10 image based on Debian Bullseye as the base image
FROM python:3.10-slim-bullseye

# Update package lists and install necessary packages
RUN apt-get update \
  && apt-get install -y --no-install-recommends --no-install-suggests \
  build-essential ffmpeg \  # Install build tools and ffmpeg (for video processing)
  && pip install --no-cache-dir --upgrade pip  # Upgrade pip

# Set the working directory inside the container to /app
WORKDIR /app

# Copy requirements.txt (Python package dependencies) to the /app directory in the container
COPY ./requirements.txt /app

# Install Python packages listed in requirements.txt
RUN pip install --no-cache-dir --requirement /app/requirements.txt

# Copy all files from the current directory (including consumer.py) to /app in the container
COPY . /app

# Define the default command to run when the container starts
CMD ["python3", "consumer.py"]
