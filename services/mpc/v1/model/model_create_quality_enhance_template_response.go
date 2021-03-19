package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type CreateQualityEnhanceTemplateResponse struct {
	// 模板ID。

	TemplateId     *int32 `json:"template_id,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o CreateQualityEnhanceTemplateResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateQualityEnhanceTemplateResponse struct{}"
	}

	return strings.Join([]string{"CreateQualityEnhanceTemplateResponse", string(data)}, " ")
}
