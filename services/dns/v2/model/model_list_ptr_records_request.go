package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListPtrRecordsRequest struct {
	Marker *string `json:"marker,omitempty"`

	Limit *int32 `json:"limit,omitempty"`

	Offset *int32 `json:"offset,omitempty"`

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	Tags *string `json:"tags,omitempty"`

	Status *string `json:"status,omitempty"`
}

func (o ListPtrRecordsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListPtrRecordsRequest struct{}"
	}

	return strings.Join([]string{"ListPtrRecordsRequest", string(data)}, " ")
}
