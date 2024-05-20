# Use the official Golang image as a parent image.
FROM golang:1.22.3

# Set the working directory inside the container.
WORKDIR /app

# Copy the current directory contents into the container at /app.
COPY . .

# Build the Go app.
RUN go build -o goreddit

# Run the compiled binary.
CMD ["./goreddit"]
