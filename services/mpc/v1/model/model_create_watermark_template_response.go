package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type CreateWatermarkTemplateResponse struct {
	// 水印模板Id

	TemplateId     *int32 `json:"template_id,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o CreateWatermarkTemplateResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateWatermarkTemplateResponse struct{}"
	}

	return strings.Join([]string{"CreateWatermarkTemplateResponse", string(data)}, " ")
}
