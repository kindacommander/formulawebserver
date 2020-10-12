package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"webserver/src/formulas"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello!")
	})
	http.HandleFunc("/lissajous", lissajousHandler)
	http.HandleFunc("/surface", surfaceHandler)
	http.HandleFunc("/mandelbrot", mandelbrotHandler)

	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func lissajousHandler(w http.ResponseWriter, r *http.Request) {
	cyclesStr := r.FormValue("cycles")
	cycles, err := strconv.Atoi(cyclesStr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "err: %v", err)
		os.Exit(1)
	}
	fmt.Fprintf(os.Stdout, "cycles: %v\n", cycles)

	formulas.Lissajous(w, cycles)
}

func surfaceHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")

	formulas.Surface(w)
}

func mandelbrotHandler(w http.ResponseWriter, r *http.Request) {
	formulas.Mandelbrot(w)
}
