package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListPrivateZonesRequest struct {
	Type string `json:"type"`

	Limit *int32 `json:"limit,omitempty"`

	Marker *string `json:"marker,omitempty"`

	Offset *int32 `json:"offset,omitempty"`

	Tags *string `json:"tags,omitempty"`

	Name *string `json:"name,omitempty"`

	Status *string `json:"status,omitempty"`

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
}

func (o ListPrivateZonesRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListPrivateZonesRequest struct{}"
	}

	return strings.Join([]string{"ListPrivateZonesRequest", string(data)}, " ")
}
