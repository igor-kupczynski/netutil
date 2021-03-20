FROM golang:1.16.2
WORKDIR /src
COPY . .
RUN go build -v -o . ./...

FROM ubuntu:20.04
LABEL org.opencontainers.image.source=https://github.com/igor-kupczynski/netutil
WORKDIR /app
COPY --from=0 /src/netutil-serve /app/netutil-serve
ENTRYPOINT ["/app/netutil-serve"]
