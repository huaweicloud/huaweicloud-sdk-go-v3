package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type AssociateSignatureKeyV2Request struct {
	// 实例编号

	InstanceId string `json:"instance_id"`

	Body *SignBindingReq `json:"body,omitempty"`
}

func (o AssociateSignatureKeyV2Request) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "AssociateSignatureKeyV2Request struct{}"
	}

	return strings.Join([]string{"AssociateSignatureKeyV2Request", string(data)}, " ")
}
