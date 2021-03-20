package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeleteAddonInstanceRequest struct {
	Id string `json:"id"`

	ClusterId string `json:"cluster_id"`
}

func (o DeleteAddonInstanceRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteAddonInstanceRequest struct{}"
	}

	return strings.Join([]string{"DeleteAddonInstanceRequest", string(data)}, " ")
}
