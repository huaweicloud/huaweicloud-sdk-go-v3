package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type UpdateBaremetalServerMetadataRequest struct {
	// 裸金属服务器ID

	ServerId string `json:"server_id"`

	Body *MetaData `json:"body,omitempty"`
}

func (o UpdateBaremetalServerMetadataRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateBaremetalServerMetadataRequest struct{}"
	}

	return strings.Join([]string{"UpdateBaremetalServerMetadataRequest", string(data)}, " ")
}
