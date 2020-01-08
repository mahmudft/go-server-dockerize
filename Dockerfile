FROM golang:alpine
ADD . /app
WORKDIR /app
ENV PORT=2030
CMD["go", "run", "server.go"]