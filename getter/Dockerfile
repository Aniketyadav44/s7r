FROM golang:1.24

WORKDIR /app

COPY . .

RUN go mod download
RUN go build -o getter ./cmd/server/main.go

EXPOSE 5002

CMD [ "./getter" ]