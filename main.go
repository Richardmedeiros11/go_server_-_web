package main

import (
	"net/http"

	_ "github.com/Richardmedeiros11/go_server_-_web/db"
	_ "github.com/lib/pq"
)

func main() {
	http.ListenAndServe(":8000", nil)
}
