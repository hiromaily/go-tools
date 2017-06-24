package main

import (
	"fmt"
	"os"

	ck "github.com/hiromaily/golibs/web/cookie"
)

func usage() {
	fmt.Fprintf(os.Stderr, "usage: %s [domain]\n", os.Args[0])
	os.Exit(2)
}

func main() {
	if len(os.Args) != 2 {
		usage()
	}

	domain := os.Args[1]
	ck.PrintCookies(domain)
}