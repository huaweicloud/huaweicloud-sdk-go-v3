package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateRecordSetWithLineRequest struct {
	ZoneId string                      `json:"zone_id"`
	Body   *CreateRecordSetWithLineReq `json:"body,omitempty"`
}

func (o CreateRecordSetWithLineRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateRecordSetWithLineRequest struct{}"
	}

	return strings.Join([]string{"CreateRecordSetWithLineRequest", string(data)}, " ")
}
