package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateWatermarkTemplateRequest struct {
	Body *CreateWatermarkTemplateReq `json:"body,omitempty"`
}

func (o CreateWatermarkTemplateRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateWatermarkTemplateRequest struct{}"
	}

	return strings.Join([]string{"CreateWatermarkTemplateRequest", string(data)}, " ")
}
