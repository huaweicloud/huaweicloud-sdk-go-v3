package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateRecordSetRequest struct {
	// 所属zone的ID。

	ZoneId string `json:"zone_id"`

	Body *CreateRecordSetReq `json:"body,omitempty"`
}

func (o CreateRecordSetRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateRecordSetRequest struct{}"
	}

	return strings.Join([]string{"CreateRecordSetRequest", string(data)}, " ")
}
