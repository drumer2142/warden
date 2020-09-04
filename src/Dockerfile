FROM golang:alpine

RUN mkdir -p /app

WORKDIR /app

COPY ./src /app

COPY ./src/go.mod ./
COPY ./src/go.sum ./

RUN go mod download

RUN go get -u github.com/go-sql-driver/mysql

RUN CGO_ENABLED=0 GOOS=linux go build -o app .

RUN ls
RUN pwd

#RUN chmod +x wait-for-it.sh

CMD ["./app"]
