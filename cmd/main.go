package main

import ftoda "github.com/henrikkorsgaard/bff-ftoda/internal/ftoda"

func main() {
	service := ftoda.NewFTODAService()
	//how do we fuzzy search in this?
	service.GetLovforslagById(101403)
}

/**

- internal/domain -> translate into application specific constructs (vote that combines sag+aktor)
- internal/ftoda -> interface repo (raw endpoints, this should be generated)

*/
