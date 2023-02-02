FROM golang:1.18.5

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./
COPY .env ./
COPY /models ./models
COPY /storages ./storages

RUN go build -o /iris-psql

VOLUME /app

EXPOSE 4000

CMD [ "/iris-psql" ]

