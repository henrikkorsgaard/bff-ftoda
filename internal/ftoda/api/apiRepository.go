package ftoda

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"

	"github.com/henrikkorsgaard/bff-ftoda/internal/ftoda"
)

type FTODAAAPIRepository struct {
	host string
}

func NewFTODAAAPIRepository(host string) *FTODAAAPIRepository {

	return &FTODAAAPIRepository{
		host: host,
	}
}

func (repo *FTODAAAPIRepository) GetSag(id int) (ftoda.Sag, error) {

	var sag ftoda.Sag

	baseUrl, err := url.Parse("https://" + repo.host + "/api/Sag")
	if err != nil {
		fmt.Printf("error with baseurl: %s\n", err)
		return sag, err
	}
	params := url.Values{}
	params.Add("$filter", "id eq "+strconv.Itoa(id))
	params.Add("$format", "json")

	baseUrl.RawQuery = params.Encode()

	res, err := http.Get(baseUrl.String())
	if err != nil {
		fmt.Printf("error making http request: %s\n", err)
		return sag, err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("error reading the body of the request: %s\n", err)
		return sag, err
	}

	err = ParseOdata(body, &sag)
	if err != nil {
		return sag, err
	}

	fmt.Printf("%+v\n", sag)

	return sag, nil

}
