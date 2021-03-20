package quotes

import "math/rand"

// Quote represents a sourced quote
type Quote struct {
	Text   string
	Source string
}

// RandomQuote returns a random quote from the given slice
func RandomQuote(quotes []Quote) Quote {
	idx := rand.Intn(len(quotes))
	return quotes[idx]
}
