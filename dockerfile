FROM golang:1.9

WORKDIR /app
COPY . .
RUN go build -ldflags "-linkmode external -extldflags -client" -a main.go

FROM scratch
COPY --from=0 /app/main /main
CMD ["/main"]