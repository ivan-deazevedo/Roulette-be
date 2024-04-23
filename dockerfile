FROM golang:latest

# Set destination for COPY
WORKDIR /app
COPY . .

# Download Go modules
ADD . /app/

EXPOSE 4000

# Run
CMD ["go", "run", "main.go"]