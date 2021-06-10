package model

import (
	"encoding/json"

	"strings"
)

type ImStatusV2 struct {
	// 状态

	Status *int32 `json:"status,omitempty"`
	// 工单id

	IncidentId *string `json:"incident_id,omitempty"`
}

func (o ImStatusV2) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ImStatusV2 struct{}"
	}

	return strings.Join([]string{"ImStatusV2", string(data)}, " ")
}
