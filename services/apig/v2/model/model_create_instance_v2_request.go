package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateInstanceV2Request struct {
	Body *InstanceCreateReq `json:"body,omitempty"`
}

func (o CreateInstanceV2Request) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateInstanceV2Request struct{}"
	}

	return strings.Join([]string{"CreateInstanceV2Request", string(data)}, " ")
}
