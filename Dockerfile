FROM golang:1.24
LABEL maintainer="Dan Fedele <dan.fedele@gmail.com>"

WORKDIR /usr/src/app

COPY go.mod .
RUN go mod download && go mod verify


COPY . .
RUN go build -v -o /usr/local/bin/app

RUN apt-get update && apt-get install -y inetutils-ping

CMD ["/usr/local/bin/app"]
