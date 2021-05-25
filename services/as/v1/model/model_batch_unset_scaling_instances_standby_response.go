package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type BatchUnsetScalingInstancesStandbyResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchUnsetScalingInstancesStandbyResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "BatchUnsetScalingInstancesStandbyResponse struct{}"
	}

	return strings.Join([]string{"BatchUnsetScalingInstancesStandbyResponse", string(data)}, " ")
}
