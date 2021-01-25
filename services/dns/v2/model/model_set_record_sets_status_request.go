package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type SetRecordSetsStatusRequest struct {
	RecordsetId string                  `json:"recordset_id"`
	Body        *SetRecordSetsStatusReq `json:"body,omitempty"`
}

func (o SetRecordSetsStatusRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "SetRecordSetsStatusRequest struct{}"
	}

	return strings.Join([]string{"SetRecordSetsStatusRequest", string(data)}, " ")
}
