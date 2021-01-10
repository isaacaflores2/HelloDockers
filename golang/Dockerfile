FROM golang:1.12.0-alpine3.9

RUN mkdir /app

ADD hello.txt /app/hello.txt

ADD ./golang /app

WORKDIR /app

RUN go build -o main .

CMD ["/app/main", "hello.txt"]