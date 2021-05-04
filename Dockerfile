FROM golang:1.16.3

ENV GOPATH=/
ENV GO111MODULE=on

COPY . .


RUN apt-get update
RUN apt-get -y install postgresql-client


RUN go build -o app-constructor-backend ./cmd/main.go

CMD ["./app-constructor-backend"]

