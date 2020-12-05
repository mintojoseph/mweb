FROM golang:1.8 as test
WORKDIR /go/src/app
COPY . /go/src/app
RUN go test ./internal -timeout 30s

FROM golang:1.8 as build
WORKDIR /go/src/app
COPY . /go/src/app
RUN make build

FROM golang:1.8
COPY --from=build /go/src/app/mweb .
CMD ["./mweb"]