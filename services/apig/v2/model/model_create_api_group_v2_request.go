package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateApiGroupV2Request struct {
	InstanceId string `json:"instance_id"`

	Body *ApiGroupReq `json:"body,omitempty"`
}

func (o CreateApiGroupV2Request) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateApiGroupV2Request struct{}"
	}

	return strings.Join([]string{"CreateApiGroupV2Request", string(data)}, " ")
}
