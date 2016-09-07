package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	var (
		dirname, portnum string
		verbose          bool
	)

	const (
		dirmsg  = "directory from which you want to serve files"
		portmsg = "port you want to bind the server to"
		vmsg    = "enable verbosity and diplay messages"
	)

	flag.StringVar(&dirname, "dir", "", dirmsg)
	flag.StringVar(&portnum, "port", "8080", portmsg)
	flag.BoolVar(&verbose, "verbose", false, vmsg)

	flag.Parse()
	if flag.NFlag() == 0 {
		fmt.Fprintf(os.Stderr, "error : no arguments options specified\n")
	}

	_, err := os.Stat(dirname)
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		return
	}

	port := ":" + portnum
	if verbose {
		fmt.Printf("Fileserver started on port %s and directory :\n%s\n", port, dirname)
	}

	log.Fatal(http.ListenAndServe(port, http.FileServer(http.Dir(dirname))))
}
