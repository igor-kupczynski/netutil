FROM golang:1.16.2
WORKDIR /src
COPY . .
RUN go build -v -o . ./...

FROM ubuntu:20.04
LABEL org.opencontainers.image.source=https://github.com/igor-kupczynski/netutil
LABEL org.opencontainers.image.authors=https://github.com/igor-kupczynski
WORKDIR /app
COPY --from=0 /src/netutil-serve /app/netutil-serve
RUN apt-get -qq update && \
	apt-get -qq install iproute2  # for ip command
ENTRYPOINT ["/app/netutil-serve"]
