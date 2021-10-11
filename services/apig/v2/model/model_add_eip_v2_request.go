package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type AddEipV2Request struct {
	// 实例编号

	InstanceId string `json:"instance_id"`

	Body *EipBindReq `json:"body,omitempty"`
}

func (o AddEipV2Request) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "AddEipV2Request struct{}"
	}

	return strings.Join([]string{"AddEipV2Request", string(data)}, " ")
}
