package main

import(
	"fmt"
	"log"
	"net/http"
	"github.com/uraj323/bookstore_apis_mongo/pkg/routers/routes"
)

func main() {
	router := routes.SetupRoutes()

	port := 8080
	fmt.Printf("Server is running on port %d..." , port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%d", port), router))
}
