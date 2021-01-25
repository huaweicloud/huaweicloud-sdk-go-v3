package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type DeleteApiV2Response struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteApiV2Response) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteApiV2Response struct{}"
	}

	return strings.Join([]string{"DeleteApiV2Response", string(data)}, " ")
}
