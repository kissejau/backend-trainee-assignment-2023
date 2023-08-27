package main

import (
	"github.com/kissejau/backend-trainee-assignment-2023/internal/server"
	_ "github.com/lib/pq"
)

func main() {
	server := server.NewServer()
	server.Run()
}
