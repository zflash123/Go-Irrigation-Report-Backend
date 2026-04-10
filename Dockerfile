FROM golang:bookworm

WORKDIR /usr/src/app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /gobinary

EXPOSE 8080

CMD ["/gobinary"]