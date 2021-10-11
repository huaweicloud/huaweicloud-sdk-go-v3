package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type DeleteApiByVersionIdV2Response struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteApiByVersionIdV2Response) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteApiByVersionIdV2Response struct{}"
	}

	return strings.Join([]string{"DeleteApiByVersionIdV2Response", string(data)}, " ")
}
