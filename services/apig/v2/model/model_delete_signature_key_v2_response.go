package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type DeleteSignatureKeyV2Response struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteSignatureKeyV2Response) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteSignatureKeyV2Response struct{}"
	}

	return strings.Join([]string{"DeleteSignatureKeyV2Response", string(data)}, " ")
}
