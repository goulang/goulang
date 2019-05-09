FROM golang:1.11.1-stretch

WORKDIR /go/src/github.com/goulang/goulang

COPY . .

RUN go build -o main main.go

EXPOSE 8080

CMD ["./main"]
