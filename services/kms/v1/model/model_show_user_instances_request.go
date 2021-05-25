package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowUserInstancesRequest struct {
	// API版本号

	VersionId string `json:"version_id"`
}

func (o ShowUserInstancesRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowUserInstancesRequest struct{}"
	}

	return strings.Join([]string{"ShowUserInstancesRequest", string(data)}, " ")
}
