FROM golang:latest

RUN go version
ENV GOPATH=/

COPY ./ ./

# install psql
RUN apt-get update
RUN apt-get -y install postgresql-client

# build go app
RUN go mod download
RUN go build -o user-balance ./cmd/main.go

CMD ["./user-balance"]