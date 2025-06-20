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

func (s *FTODAService) GetLovforslagById(id int) (sag Sag, err error) {
	//First we should check a database, but that is not created yet
	//If not found in database, then we get it from the api

	query := odataQuery{
		entity: "Sag",
		filter: "typeid eq 3 and id eq " + strconv.Itoa(id),
	}

	// this need to be moved into a different repo service
	odata, err := s.api.getData(query)
	if err != nil {
		fmt.Printf("error from getLovforslag: %s\n", err)
		return sag, err
	}

	// this need to be moved into a different repo service
	var sager []Sag
	err = json.Unmarshal(odata.Result, &sager)
	if err != nil {
		fmt.Printf("error from getLovforslag: %s\n", err)
		return sag, err
	}

	return sager[0], nil
}

// offset map into skip next for now
func (s *FTODAService) GetLovforslag(offset int) ([]Sag, error) {
	query := odataQuery{
		entity: "Sag",
		filter: "typeid eq 3",
		skip:   offset,
	}

	odata, err := s.api.getData(query)
	if err != nil {
		fmt.Printf("error from getLovforslag: %s\n", err)
		return nil, err
	}

	var sager []Sag
	err = json.Unmarshal(odata.Result, &sager)
	if err != nil {
		fmt.Printf("error from getLovforslag: %s\n", err)
		return nil, err
	}

	return sager, nil
}
