FROM golang:1.22.5-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY src/ ./src/

# RUN go build main.go
RUN go build src/main.go

EXPOSE 9096

CMD [ "./main" ]