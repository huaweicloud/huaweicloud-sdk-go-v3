package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type UpdateApiV2Request struct {
	InstanceId string     `json:"instance_id"`
	ApiId      string     `json:"api_id"`
	Body       *ApiCreate `json:"body,omitempty"`
}

func (o UpdateApiV2Request) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateApiV2Request struct{}"
	}

	return strings.Join([]string{"UpdateApiV2Request", string(data)}, " ")
}
