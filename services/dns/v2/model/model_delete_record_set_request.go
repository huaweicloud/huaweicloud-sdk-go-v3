package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeleteRecordSetRequest struct {
	ZoneId string `json:"zone_id"`

	RecordsetId string `json:"recordset_id"`
}

func (o DeleteRecordSetRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteRecordSetRequest struct{}"
	}

	return strings.Join([]string{"DeleteRecordSetRequest", string(data)}, " ")
}
