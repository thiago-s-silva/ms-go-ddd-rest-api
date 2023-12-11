# Use the official Golang image with a specific version
FROM golang:alpine3.19

# Define the port placeholder
ARG PORT=8080

# Set the working directory
WORKDIR /app

# Move project files
COPY . .

# Define the exposed port based on the environment variable
EXPOSE ${PORT}

# Build app
RUN CGO_ENABLED=0 GOOS=linux go build -o main cmd/main.go

# Run the application
CMD ["./main"]
