package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3333"
	}
	ns := os.Getenv("NAMESPACE")

	handler := http.NewServeMux()
	handler.HandleFunc("/dapr/config", func(w http.ResponseWriter, r *http.Request) {
		start, err := strconv.Atoi(port)
		if err != nil {
			start = 0
		}
		start = start * 100 * rand.Intn(10)
		numTypes := 10

		types := []string{}
		for k := start; k < start+numTypes; k++ {
			types = append(types, ns+"-"+port+"-myactortype-"+strconv.Itoa(k))
		}

		w.Write([]byte(fmt.Sprintf(`{"entities": ["%s"]}`, strings.Join(types, `","`))))
	})
	handler.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`OK`))
	})

	handler.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`OK`))
	})
	handler.HandleFunc("/healthz/outbound", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`OK`))
	})
	err := http.ListenAndServe(fmt.Sprintf(":%s", port), handler)
	if err != nil {
		panic(err)
	}
}
