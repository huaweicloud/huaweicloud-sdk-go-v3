package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeleteRecordSetsRequest struct {
	ZoneId      string `json:"zone_id"`
	RecordsetId string `json:"recordset_id"`
}

func (o DeleteRecordSetsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteRecordSetsRequest struct{}"
	}

	return strings.Join([]string{"DeleteRecordSetsRequest", string(data)}, " ")
}
