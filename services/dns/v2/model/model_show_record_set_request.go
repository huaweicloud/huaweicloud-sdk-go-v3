package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowRecordSetRequest struct {
	ZoneId string `json:"zone_id"`

	RecordsetId string `json:"recordset_id"`
}

func (o ShowRecordSetRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowRecordSetRequest struct{}"
	}

	return strings.Join([]string{"ShowRecordSetRequest", string(data)}, " ")
}
