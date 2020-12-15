FROM golang:latest

WORKDIR /app/statGet/cmd
COPY cmd .

WORKDIR /app/statGet
COPY main.go .env go.mod ./

RUN go mod download
RUN go build .

RUN chmod +x ./statGet

CMD ["/app/statGet/statGet"]