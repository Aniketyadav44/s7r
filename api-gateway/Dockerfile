FROM golang:1.24

WORKDIR /app

COPY . .

RUN go mod download
RUN go build -o api-gateway ./cmd/api/main.go

EXPOSE 8000

CMD [ "./api-gateway" ]