# Use an official Golang runtime as a parent image
FROM golang:latest

# Set the working directory to /app
WORKDIR /app

# Copy the current directory contents into the container at /app
COPY . /app

# update sh file
RUN mv start.sh.sample start.sh

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["sh", "start.sh"]
