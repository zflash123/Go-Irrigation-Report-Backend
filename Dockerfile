FROM golang:bookworm

WORKDIR /usr/src/app

ARG DB_HOST
ARG DB_PORT
ARG DB_USER
ARG DB_PWD
ARG DB_NAME

RUN echo 'DB_HOST="'$DB_HOST'"' >> .env
RUN echo 'DB_PORT="'$DB_PORT'"' >> .env
RUN echo 'DB_USER="'$DB_USER'"' >> .env
RUN echo 'DB_PWD="'$DB_PWD'"' >> .env
RUN echo 'DB_NAME="'$DB_NAME'"' >> .env

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /gobinary

EXPOSE 8080

CMD ["/gobinary"]