package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeleteApiV2Request struct {
	ProjectId  string `json:"project_id"`
	InstanceId string `json:"instance_id"`
	ApiId      string `json:"api_id"`
}

func (o DeleteApiV2Request) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteApiV2Request struct{}"
	}

	return strings.Join([]string{"DeleteApiV2Request", string(data)}, " ")
}
