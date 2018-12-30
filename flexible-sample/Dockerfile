FROM golang:1.11

ADD . /go/src/flexible-sample

WORKDIR /go/src/flexible-sample

EXPOSE 8080

ENV GO111MODULE=on

CMD ["go", "run", "web/main.go"]
