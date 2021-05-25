package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type BatchProtectScalingInstancesResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchProtectScalingInstancesResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "BatchProtectScalingInstancesResponse struct{}"
	}

	return strings.Join([]string{"BatchProtectScalingInstancesResponse", string(data)}, " ")
}
