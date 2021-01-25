package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListWatermarkTemplateRequest struct {
	TemplateId *[]int32 `json:"template_id,omitempty"`
	Page       *int32   `json:"page,omitempty"`
	Size       *int32   `json:"size,omitempty"`
}

func (o ListWatermarkTemplateRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListWatermarkTemplateRequest struct{}"
	}

	return strings.Join([]string{"ListWatermarkTemplateRequest", string(data)}, " ")
}
