package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListCustomLineRequest struct {
	LineId *string `json:"line_id,omitempty"`

	Name *string `json:"name,omitempty"`

	Limit *int32 `json:"limit,omitempty"`

	Offset *int32 `json:"offset,omitempty"`

	ShowDetail *bool `json:"show_detail,omitempty"`
}

func (o ListCustomLineRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListCustomLineRequest struct{}"
	}

	return strings.Join([]string{"ListCustomLineRequest", string(data)}, " ")
}
