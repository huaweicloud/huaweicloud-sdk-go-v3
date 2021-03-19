package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeleteSignatureKeyV2Request struct {
	InstanceId string `json:"instance_id"`

	SignId string `json:"sign_id"`
}

func (o DeleteSignatureKeyV2Request) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteSignatureKeyV2Request struct{}"
	}

	return strings.Join([]string{"DeleteSignatureKeyV2Request", string(data)}, " ")
}
