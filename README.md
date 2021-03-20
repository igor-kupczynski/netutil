# netutil

## Serve

Serves random quotes over http.

### Usage

Run as a docker container:

    $ docker run --rm -ti -p 8080:8080 --name netutil-serve ghcr.io/igor-kupczynski/netutil:latest

Or as a binary:

    $ go get github.com/igor-kupczynski/netutil/cmd/netutil-serve
    $ netutil-serve

Then in a different terminal:

    $ curl localhost:8080/quote
    Program testing can be used to show the presence of bugs, but never to show their absence!
    
      -- Dijkstra (1970) "Notes On Structured Programming" (EWD249), Section 3 ("On The Reliability of Mechanisms"), corollary at the end.


### Build from scratch

    $ make
    $ ./netutil-serve -h


Or in a docker container: 

    docker build -t ghcr.io/igor-kupczynski/netutil:dev .
