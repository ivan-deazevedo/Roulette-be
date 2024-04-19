FROM golang:1.22

# Set destination for COPY
WORKDIR /app

# Download Go modules
ADD . /app/

EXPOSE 4000

# Run
CMD ["go", "run", "main.go"]