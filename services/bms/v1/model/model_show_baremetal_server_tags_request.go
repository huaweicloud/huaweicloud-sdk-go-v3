package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowBaremetalServerTagsRequest struct {
	ServerId string `json:"server_id"`
}

func (o ShowBaremetalServerTagsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowBaremetalServerTagsRequest struct{}"
	}

	return strings.Join([]string{"ShowBaremetalServerTagsRequest", string(data)}, " ")
}
