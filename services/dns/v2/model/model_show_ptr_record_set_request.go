package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowPtrRecordSetRequest struct {
	Region string `json:"region"`

	FloatingipId string `json:"floatingip_id"`
}

func (o ShowPtrRecordSetRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowPtrRecordSetRequest struct{}"
	}

	return strings.Join([]string{"ShowPtrRecordSetRequest", string(data)}, " ")
}
