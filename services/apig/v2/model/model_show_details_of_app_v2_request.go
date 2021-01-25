package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowDetailsOfAppV2Request struct {
	ProjectId  string `json:"project_id"`
	InstanceId string `json:"instance_id"`
	AppId      string `json:"app_id"`
}

func (o ShowDetailsOfAppV2Request) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowDetailsOfAppV2Request struct{}"
	}

	return strings.Join([]string{"ShowDetailsOfAppV2Request", string(data)}, " ")
}
