package main

import ftoda "github.com/henrikkorsgaard/bff-ftoda/internal/ftoda/api"

func main() {
	repo := ftoda.NewFTODAAAPIRepository("oda.ft.dk")
	repo.GetSag(10)
}

/**

- internal/domain -> translate into application specific constructs (vote that combines sag+aktor)
- internal/ftoda -> interface repo (raw endpoints, this should be generated)

*/
