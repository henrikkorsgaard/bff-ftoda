package ftoda

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strings"
)

type Odata struct {
	Metadata string            `json:"odata.metadata"`
	Result   []json.RawMessage `json:"value"`
}

// TODO: Change to handle array OR create version that always just returns one

func ParseOdata(data []byte, v any) error {

	var odata Odata
	err := json.Unmarshal(data, &odata)
	if err != nil {
		fmt.Printf("Unable to marshal JSON: %s", err)
		return err
	}

	tag, err := extractType(odata.Metadata)
	if err != nil {
		return err
	}

	// breaks if the package breaks
	//not sure we need this

	if "*ftoda."+tag == reflect.TypeOf(v).String() {
		fmt.Println("This is correct, we do not need to run amok")
		err := json.Unmarshal(odata.Result[0], v)
		if err != nil {
			fmt.Println(err)
			return err
		}
	}

	return nil

}

func extractType(metadata string) (string, error) {
	splits := strings.SplitAfter(metadata, "#")
	if len(splits) > 0 {
		return splits[1], nil
	} else {
		return "", errors.New("Unable to detect struct tag for " + metadata)
	}
}
