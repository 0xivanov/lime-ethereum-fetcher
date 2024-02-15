# Use an official Golang runtime as a parent 
FROM golang:1.20.3-alpine

# Set destination for COPY
WORKDIR /app

COPY . .

# Download Go modules
RUN go mod download

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /lime-api

EXPOSE 9090

CMD [ "/lime-api" ]
