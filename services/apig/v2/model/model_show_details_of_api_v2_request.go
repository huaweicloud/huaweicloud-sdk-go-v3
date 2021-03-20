package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowDetailsOfApiV2Request struct {
	InstanceId string `json:"instance_id"`

	ApiId string `json:"api_id"`
}

func (o ShowDetailsOfApiV2Request) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowDetailsOfApiV2Request struct{}"
	}

	return strings.Join([]string{"ShowDetailsOfApiV2Request", string(data)}, " ")
}
