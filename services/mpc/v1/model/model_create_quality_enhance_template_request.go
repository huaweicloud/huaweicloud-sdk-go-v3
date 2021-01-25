package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateQualityEnhanceTemplateRequest struct {
	Body *QualityEnhanceTemplate `json:"body,omitempty"`
}

func (o CreateQualityEnhanceTemplateRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateQualityEnhanceTemplateRequest struct{}"
	}

	return strings.Join([]string{"CreateQualityEnhanceTemplateRequest", string(data)}, " ")
}
