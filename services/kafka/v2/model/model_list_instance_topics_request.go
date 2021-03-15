package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListInstanceTopicsRequest struct {
	InstanceId string `json:"instance_id"`
}

func (o ListInstanceTopicsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListInstanceTopicsRequest struct{}"
	}

	return strings.Join([]string{"ListInstanceTopicsRequest", string(data)}, " ")
}
