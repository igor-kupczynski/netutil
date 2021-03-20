package main

import (
	"flag"
	"fmt"
	"github.com/igor-kupczynski/netutil/quotes"
	"log"
	"net/http"
	"os"
)

var portFlag = flag.Int("port", 8080, "port to run the server on")

func main() {
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
	}

	flag.Parse()

	if err := run(*portFlag); err != nil {
		log.Fatal(err)
	}
}

func run(port int) error {
	http.HandleFunc("/quote", func(w http.ResponseWriter, r *http.Request) {
		q := quotes.RandomQuote(quotes.Dijkstra)
		resp := fmt.Sprintf("%s\n\n  -- %s\n", q.Text, q.Source)
		fmt.Fprintf(w, resp)
	})

	addr := fmt.Sprintf(":%d", port)
	log.Printf("Listening on %s\n", addr)
	return http.ListenAndServe(addr, nil)
}
