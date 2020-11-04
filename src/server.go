package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"os"
	"strconv"
	"webserver/src/eval"
	"webserver/src/formulas"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello!")
	})
	http.HandleFunc("/lissajous", lissajousHandler)
	http.HandleFunc("/plot", plot)
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

func mandelbrotHandler(w http.ResponseWriter, r *http.Request) {
	formulas.Mandelbrot(w)
}

func plot(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	expr, err := formulas.ParseAndCheck(r.Form.Get("expr"))
	if err != nil {
		http.Error(w, "unknown expression: "+err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "image/svg+xml")

	f := func(x, y float64) float64 {
		r := math.Hypot(x, y)
		return expr.Eval(eval.Env{"x": x, "y": y, "r": r})
	}

	formulas.Surface(w, f)
}
