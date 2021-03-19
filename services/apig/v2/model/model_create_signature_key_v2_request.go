package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateSignatureKeyV2Request struct {
	InstanceId string `json:"instance_id"`

	Body *SignatureReq `json:"body,omitempty"`
}

func (o CreateSignatureKeyV2Request) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateSignatureKeyV2Request struct{}"
	}

	return strings.Join([]string{"CreateSignatureKeyV2Request", string(data)}, " ")
}
