FROM golang:1.8 as test
WORKDIR /go/src/app
COPY . /go/src/app
RUN go test -timeout 30s

FROM golang:1.8 as build
WORKDIR /go/src/app
COPY . /go/src/app
RUN make build
COPY . /bin
CMD ["/bin/mweb"]