FROM golang:1.17.0-alpine3.14 as builder
COPY server /code/server
WORKDIR /code/server

RUN go get github.com/gin-gonic/gin
RUN go get github.com/gin-gonic/gin/binding@v1.7.4
RUN go get github.com/go-sql-driver/mysql
RUN go get github.com/jinzhu/gorm
RUN go get go.uber.org/zap
RUN go get -u github.com/cosmtrek/air
# local debug
RUN go install github.com/go-delve/delve/cmd/dlv@latest

#CMD ["go", "run", "main.go"]
CMD ["air", "-c", ".air.toml"]