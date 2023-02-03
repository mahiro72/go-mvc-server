FROM golang:1.19 AS dev

WORKDIR /go/src/github.com/mahiro72/go-mvc-server

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

RUN go install github.com/cosmtrek/air@v1.29.0

CMD ["air", "-c", ".air.toml"]
