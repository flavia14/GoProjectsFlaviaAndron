package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Helo Flavia!"))
	})

	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		panic(err.Error())
	}
}
