package main

import (
	"fmt"
	"log"

	"github.com/edwintantawi/udoit/internal/server"
	"github.com/edwintantawi/udoit/pkg/sqlite"
)

const addr = ":5000"

func main() {
	db := sqlite.New("data.db")
	defer db.Close()

	srv := server.NewHTTP(db)
	fmt.Printf("Server running at %s", addr)
	log.Fatal(srv.Setup(addr).ListenAndServe())
}
