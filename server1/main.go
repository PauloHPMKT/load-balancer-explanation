package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// This is to simulate a slow response time to test least connections algorithm LB
		time.Sleep(15 * time.Second)
		fmt.Fprintf(w, "Hello, World! from server 1")
	})

	log.Fatal(http.ListenAndServe(":8081", nil))
}

/**
To simulate a load on the server, you can use the following command:
for i in 1 2 3 4 5 6 7 8 9; do
  curl -s "http://localhost:8080/" &
done
wait
*/
