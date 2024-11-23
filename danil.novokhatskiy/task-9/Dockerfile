FROM golang:1.23

WORKDIR /task-9

COPY . .

RUN go get -v ./...

RUN go build -o api ./cmd/service/main.go

EXPOSE 8000

CMD ["./api"]