package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type DeleteScalingTagInfoResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteScalingTagInfoResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteScalingTagInfoResponse struct{}"
	}

	return strings.Join([]string{"DeleteScalingTagInfoResponse", string(data)}, " ")
}
