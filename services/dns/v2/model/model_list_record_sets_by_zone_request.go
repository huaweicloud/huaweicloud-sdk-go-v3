package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListRecordSetsByZoneRequest struct {
	ZoneId string `json:"zone_id"`

	Marker *string `json:"marker,omitempty"`

	Limit *string `json:"limit,omitempty"`

	Offset *string `json:"offset,omitempty"`

	Tags *string `json:"tags,omitempty"`

	Status *string `json:"status,omitempty"`

	Type *string `json:"type,omitempty"`

	Name *string `json:"name,omitempty"`

	Id *string `json:"id,omitempty"`

	SortKey *string `json:"sort_key,omitempty"`

	SortDir *string `json:"sort_dir,omitempty"`
}

func (o ListRecordSetsByZoneRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListRecordSetsByZoneRequest struct{}"
	}

	return strings.Join([]string{"ListRecordSetsByZoneRequest", string(data)}, " ")
}
