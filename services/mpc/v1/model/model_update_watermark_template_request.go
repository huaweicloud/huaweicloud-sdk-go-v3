package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type UpdateWatermarkTemplateRequest struct {
	Body *WatermarkTemplate `json:"body,omitempty"`
}

func (o UpdateWatermarkTemplateRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateWatermarkTemplateRequest struct{}"
	}

	return strings.Join([]string{"UpdateWatermarkTemplateRequest", string(data)}, " ")
}
