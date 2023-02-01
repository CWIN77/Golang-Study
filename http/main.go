package main

import (
	"fmt"
	"math"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		const base = 8
		result := 0.0
		for i := math.Pow(base, 7) * 40; i >= 0; i-- {
			result += i + math.Pow(i, 10)
		}
		w.Write([]byte(fmt.Sprintf("%f", result)))
	})

	http.ListenAndServe(":3001", nil)
}
