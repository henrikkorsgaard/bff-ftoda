package ftoda

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type FTODAService struct {
	api *apiRepository
}

func NewFTODAService() FTODAService {

	// Host should come from either a factory or .env
	repo := newAPIRepository("oda.ft.dk")

	return FTODAService{
		api: repo,
	}
}

func (s *FTODAService) GetLovforslagById(id int) {
	//First we should check a database, but that is not created yet
	//If not found in database, then we get it from the api

	query := odataQuery{
		entity: "Sag",
		filter: "typeid eq 3 and id eq " + strconv.Itoa(id),
	}

	odata, err := s.api.getData(query)
	if err != nil {
		fmt.Printf("error from getLovforslag: %s\n", err)
	}

	var sager []Sag
	err = json.Unmarshal(odata.Result, &sager)
	if err != nil {
		fmt.Printf("error from getLovforslag: %s\n", err)
	}

	//fmt.Printf("%+v\n", sager)
}
