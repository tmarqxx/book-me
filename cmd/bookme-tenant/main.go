package main

import (
	"book-me/internal/database"
	"book-me/internal/transport/http"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	port := flag.String("port", "3000", "Port to bind server to")
	localDb := flag.String("local-db", "bookme-tenant.local.db", "Filename for local embedded database")
	tursoUrl := flag.String("turso-url", "", "Turso database URL")
	tursoToken := flag.String("turso-token", "", "Turso auth token")

	flag.Parse()

	if *tursoUrl == "" {
		log.Fatal("--turso-url flag is required")
	}

	if *tursoToken == "" {
		log.Fatal("--turso-token flag is required")
	}

	db, dir := database.OpenEmbeddedDB(*localDb, *tursoUrl, *tursoToken)
	fmt.Println("Hello??")
	defer db.Close()
	defer os.RemoveAll(dir)

	server := http.NewTenantServer()

	server.Logger.Fatal(server.Start(":" + *port))
}
