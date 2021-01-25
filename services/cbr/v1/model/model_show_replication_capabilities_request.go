package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowReplicationCapabilitiesRequest struct {
}

func (o ShowReplicationCapabilitiesRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowReplicationCapabilitiesRequest struct{}"
	}

	return strings.Join([]string{"ShowReplicationCapabilitiesRequest", string(data)}, " ")
}
