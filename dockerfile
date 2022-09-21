FROM golang:1.9

WORKDIR /go/src/github.com/purplebooth/example
COPY . .
RUN go build -ldflags "-linkmode external -extldflags -client" -a main.go

FROM scratch
COPY --from=0 /go/src/github.com/purplebooth/example/main /main
CMD ["/main"]