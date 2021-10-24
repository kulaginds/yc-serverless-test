FROM golang:1.16 as build

COPY . /build

WORKDIR /build

RUN go build -o hello hello.go

FROM alpine:3

COPY --from=build hello /bin

CMD ['hello']
