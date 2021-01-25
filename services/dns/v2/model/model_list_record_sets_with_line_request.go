package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListRecordSetsWithLineRequest struct {
	ZoneType      *string `json:"zone_type,omitempty"`
	Marker        *string `json:"marker,omitempty"`
	Limit         *string `json:"limit,omitempty"`
	Offset        *string `json:"offset,omitempty"`
	LineId        *string `json:"line_id,omitempty"`
	Tags          *string `json:"tags,omitempty"`
	Status        *string `json:"status,omitempty"`
	Type          *string `json:"type,omitempty"`
	Name          *string `json:"name,omitempty"`
	Id            *string `json:"id,omitempty"`
	Records       *string `json:"records,omitempty"`
	SortKey       *string `json:"sort_key,omitempty"`
	SortDir       *string `json:"sort_dir,omitempty"`
	HealthCheckId *string `json:"health_check_id,omitempty"`
	SearchMode    *string `json:"search_mode,omitempty"`
}

func (o ListRecordSetsWithLineRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListRecordSetsWithLineRequest struct{}"
	}

	return strings.Join([]string{"ListRecordSetsWithLineRequest", string(data)}, " ")
}
