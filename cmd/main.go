package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/henrikkorsgaard/bff-ftoda/handlers"
)

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/lovforslag", handlers.GetLovforslag).Methods("GET")
	r.HandleFunc("/lovforslag/{id}", handlers.GetLovforslagById).Methods("GET")
	// Start the server
	fmt.Println("Server is running on port 8000...")
	log.Fatal(http.ListenAndServe(":8000", r))
}

/**

- internal/domain -> translate into application specific constructs (vote that combines sag+aktor)
- internal/ftoda -> interface repo (raw endpoints, this should be generated)

*/
