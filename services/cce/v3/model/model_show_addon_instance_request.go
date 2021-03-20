package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowAddonInstanceRequest struct {
	Id string `json:"id"`

	ClusterId string `json:"cluster_id"`
}

func (o ShowAddonInstanceRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowAddonInstanceRequest struct{}"
	}

	return strings.Join([]string{"ShowAddonInstanceRequest", string(data)}, " ")
}
