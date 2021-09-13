FROM golang:1.17.0-alpine3.14 as builder
COPY server /code/server
WORKDIR /code/server

RUN go get -u github.com/cosmtrek/air

#CMD ["go", "run", "main.go"]
CMD ["air", "-c", ".air.toml"]