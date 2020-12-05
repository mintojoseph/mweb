FROM golang:1.9 as test
WORKDIR /go/src/app
COPY . /go/src/app
RUN go test ./internal -timeout 30s

FROM golang:1.9 as build
WORKDIR /go/src/app
COPY . /go/src/app
RUN make build

FROM scratch
COPY --from=build /go/src/app/mweb .
CMD ["./mweb"]