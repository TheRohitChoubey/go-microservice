FROM golang:1.9

WORKDIR /go/src/github.com/therohitchoubey/go-microservice
COPY . .
RUN go build -ldflags "-linkmode external -extldflags -client" -a main.go

FROM scratch
COPY --from=0 /go/src/github.com/therohitchoubey/go-microservice/main /main
CMD ["/main"]