package subhello

import "rsc.io/quote/v3"

func Subhello() string {
	return quote.HelloV3()
}

func Subpoverb() string {
	return quote.Concurrency()
}
