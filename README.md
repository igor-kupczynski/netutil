# netutil

## Serve

Serves a list of quotes over http.

### Installation

#### Install with `go get`

    $ go get github.com/igor-kupczynski/netutil/cmd/netutil-serve
    $ netutil-serve -h


#### Build from scratch

    $ make
    $ ./netutil-serve -h

#### Build from scratch using a docker container

    docker build -t ghcr.io/igor-kupczynski/netutil:dev .


### Usage

    $ netutil-serve

Or with a docker image:

    $ docker run -p 8080:8080 --name netutil-serve -ti --rm ghcr.io/igor-kupczynski/netutil:dev 


Then in a different terminal:

    $ curl localhost:8080/quote
    Program testing can be used to show the presence of bugs, but never to show their absence!
    
      -- Dijkstra (1970) "Notes On Structured Programming" (EWD249), Section 3 ("On The Reliability of Mechanisms"), corollary at the end.
