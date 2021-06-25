package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListWatermarkTemplateResponse struct {
	Templates *[]WatermarkTemplate `json:"templates,omitempty"`

	Total          *int32 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListWatermarkTemplateResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListWatermarkTemplateResponse struct{}"
	}

	return strings.Join([]string{"ListWatermarkTemplateResponse", string(data)}, " ")
}
