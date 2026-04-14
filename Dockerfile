FROM golang:bookworm

WORKDIR /usr/src/app

ARG DB_HOST
ARG DB_PORT
ARG DB_USER
ARG DB_PWD
ARG DB_NAME
ARG GCS_BUCKET
ARG JWT_KEY
ARG HTTP_REQUEST_LIMIT_IN_MB

RUN echo 'DB_HOST="'$DB_HOST'"' >> .env
RUN echo 'DB_PORT="'$DB_PORT'"' >> .env
RUN echo 'DB_USER="'$DB_USER'"' >> .env
RUN echo 'DB_PWD="'$DB_PWD'"' >> .env
RUN echo 'DB_NAME="'$DB_NAME'"' >> .env
RUN echo 'GCS_BUCKET="'$GCS_BUCKET'"' >> .env
RUN echo 'JWT_KEY="'$JWT_KEY'"' >> .env
RUN echo 'HTTP_REQUEST_LIMIT_IN_MB="'$HTTP_REQUEST_LIMIT_IN_MB'"' >> .env

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o gobinary

ENTRYPOINT ["/usr/src/app/gobinary"]