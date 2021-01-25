package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeleteAppV2Request struct {
	ProjectId  string `json:"project_id"`
	InstanceId string `json:"instance_id"`
	AppId      string `json:"app_id"`
}

func (o DeleteAppV2Request) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteAppV2Request struct{}"
	}

	return strings.Join([]string{"DeleteAppV2Request", string(data)}, " ")
}
