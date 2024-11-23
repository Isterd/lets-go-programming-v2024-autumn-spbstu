FROM golang:1.23

WORKDIR /app

COPY . .

RUN go get -d v ./...

RUN go build -o github.com/katagiriwhy/task-9 .

EXPOSE 8000

CMD ["./api"]