package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type DeleteApiGroupV2Response struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteApiGroupV2Response) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteApiGroupV2Response struct{}"
	}

	return strings.Join([]string{"DeleteApiGroupV2Response", string(data)}, " ")
}
