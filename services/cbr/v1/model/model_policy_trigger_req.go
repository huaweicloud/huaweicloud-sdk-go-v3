package model

import (
	"encoding/json"

	"strings"
)

type PolicyTriggerReq struct {
	Properties *PolicyTriggerPropertiesReq `json:"properties"`
}

func (o PolicyTriggerReq) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "PolicyTriggerReq struct{}"
	}

	return strings.Join([]string{"PolicyTriggerReq", string(data)}, " ")
}
