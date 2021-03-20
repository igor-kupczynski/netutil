# netutil

## Serve

Serves a list of quotes over http.

### Install

    $ go get github.com/igor-kupczynski/netutil/cmd/netutil-serve
    $ netutil-serve -h


### Build from scratch

    $ make
    $ ./netutil-serve -h


### Usage

    $ netutil-serve

Then in a different terminal:

    $ curl localhost:8080/quote
    Program testing can be used to show the presence of bugs, but never to show their absence!
    
      -- Dijkstra (1970) "Notes On Structured Programming" (EWD249), Section 3 ("On The Reliability of Mechanisms"), corollary at the end.
