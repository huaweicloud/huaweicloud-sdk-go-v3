package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type UpdateSignatureKeyV2Request struct {
	InstanceId string        `json:"instance_id"`
	SignId     string        `json:"sign_id"`
	Body       *SignatureReq `json:"body,omitempty"`
}

func (o UpdateSignatureKeyV2Request) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateSignatureKeyV2Request struct{}"
	}

	return strings.Join([]string{"UpdateSignatureKeyV2Request", string(data)}, " ")
}
