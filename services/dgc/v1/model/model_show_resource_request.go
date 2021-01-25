package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowResourceRequest struct {
	ResourceId string `json:"resource_id"`
}

func (o ShowResourceRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowResourceRequest struct{}"
	}

	return strings.Join([]string{"ShowResourceRequest", string(data)}, " ")
}
