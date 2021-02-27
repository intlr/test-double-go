FROM golang:1.16-alpine

WORKDIR /go/src/github.com/alr-lab/test-double-go/

COPY . /go/src/github.com/alr-lab/test-double-go/

CMD go test -v ./...
