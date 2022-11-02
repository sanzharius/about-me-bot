FROM golang:latest

ENV GOPATH /go
ENV PORT = 8080


COPY . /telegrambot/sanzhar/
WORKDIR /telegrambot/sanzhar/


RUN go mod download
RUN go build -o main .
EXPOSE 8080
COPY .env /app
CMD ["./main"]


