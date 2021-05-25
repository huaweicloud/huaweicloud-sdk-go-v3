package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type BatchSetScalingInstancesStandbyResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchSetScalingInstancesStandbyResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "BatchSetScalingInstancesStandbyResponse struct{}"
	}

	return strings.Join([]string{"BatchSetScalingInstancesStandbyResponse", string(data)}, " ")
}
