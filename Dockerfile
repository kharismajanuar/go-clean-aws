
FROM golang:1.20.4-alpine

RUN mkdir /app

WORKDIR /app

COPY ./ /app

RUN go mod tidy

RUN go build -o gocleanaws

CMD ["./gocleanaws"]