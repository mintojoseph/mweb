FROM golang:1.14-alpine as build
RUN apk add make git
WORKDIR /go/src/app
COPY . /go/src/app
RUN make build

FROM ubuntu:18.04
COPY --from=build /go/src/app/mweb /bin/mweb
CMD ["/bin/mweb"]

