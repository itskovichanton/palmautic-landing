FROM golang:1.19-alpine

WORKDIR /app

COPY . .

RUN go mod download

RUN go build cmd/main.go

EXPOSE 3003

CMD [ "./main", "config/config.yml" ]