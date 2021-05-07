package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/json", helloFunc)

	fmt.Println("Server Start!!")
	http.ListenAndServe(":8000", nil)
}

func helloFunc(w http.ResponseWriter, r *http.Request) {
	log.Print(r)

	ans := map[string]string {
		"message": "hello",
	}
	res, err := json.Marshal(ans)
	if err != nil {
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}
