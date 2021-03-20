package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListRestoreDatabasesRequest struct {
	XLanguage *string `json:"X-Language,omitempty"`

	InstanceId string `json:"instance_id"`

	RestoreTime string `json:"restore_time"`

	Offset *int32 `json:"offset,omitempty"`

	Limit *int32 `json:"limit,omitempty"`
}

func (o ListRestoreDatabasesRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListRestoreDatabasesRequest struct{}"
	}

	return strings.Join([]string{"ListRestoreDatabasesRequest", string(data)}, " ")
}
