package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type DeleteInstancesV2Response struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteInstancesV2Response) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteInstancesV2Response struct{}"
	}

	return strings.Join([]string{"DeleteInstancesV2Response", string(data)}, " ")
}
